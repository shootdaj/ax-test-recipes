//go:build scenario

package scenarios

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shootdaj/ax-test-recipes/internal/models"
	"github.com/shootdaj/ax-test-recipes/internal/router"
	"github.com/shootdaj/ax-test-recipes/internal/store"
)

func setupServer() *httptest.Server {
	s := store.New()
	handler := router.New(s)
	return httptest.NewServer(handler)
}

func createRecipe(t *testing.T, srv *httptest.Server, recipe models.Recipe) models.Recipe {
	t.Helper()
	body, _ := json.Marshal(recipe)
	resp, err := http.Post(srv.URL+"/api/recipes", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("create recipe failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}
	var created models.Recipe
	json.NewDecoder(resp.Body).Decode(&created)
	return created
}

func TestFullRecipeLifecycle(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	// 1. Create a recipe
	pancakes := createRecipe(t, srv, models.Recipe{
		Title:       "Pancakes",
		Description: "Fluffy breakfast pancakes",
		Ingredients: []models.Ingredient{
			{Name: "flour", Quantity: 2, Unit: "cups"},
			{Name: "milk", Quantity: 1.5, Unit: "cups"},
			{Name: "eggs", Quantity: 2, Unit: "pcs"},
		},
		Steps:    []string{"Mix dry ingredients", "Add wet ingredients", "Cook on griddle"},
		PrepTime: 10,
		CookTime: 15,
		Servings: 4,
		Tags:     []string{"breakfast", "quick"},
		Category: "breakfast",
	})

	// 2. Verify it appears in listing
	resp, _ := http.Get(srv.URL + "/api/recipes")
	var recipes []models.Recipe
	json.NewDecoder(resp.Body).Decode(&recipes)
	resp.Body.Close()
	if len(recipes) != 1 {
		t.Fatalf("expected 1 recipe, got %d", len(recipes))
	}

	// 3. Update the recipe
	pancakes.Title = "Fluffy Pancakes"
	body, _ := json.Marshal(pancakes)
	req, _ := http.NewRequest(http.MethodPut, srv.URL+"/api/recipes/"+pancakes.ID, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	updateResp, _ := http.DefaultClient.Do(req)
	updateResp.Body.Close()

	// 4. Search by ingredient
	searchResp, _ := http.Get(srv.URL + "/api/search?ingredient=flour")
	var searchResults []models.Recipe
	json.NewDecoder(searchResp.Body).Decode(&searchResults)
	searchResp.Body.Close()
	if len(searchResults) != 1 {
		t.Errorf("expected 1 search result, got %d", len(searchResults))
	}

	// 5. Scale the recipe
	scaleResp, _ := http.Get(srv.URL + "/api/recipes/" + pancakes.ID + "/scale?servings=8")
	var scaled models.Recipe
	json.NewDecoder(scaleResp.Body).Decode(&scaled)
	scaleResp.Body.Close()
	if scaled.Servings != 8 {
		t.Errorf("expected 8 servings, got %d", scaled.Servings)
	}
	if scaled.Ingredients[0].Quantity != 4 {
		t.Errorf("expected 4 cups flour, got %f", scaled.Ingredients[0].Quantity)
	}

	// 6. Delete the recipe
	delReq, _ := http.NewRequest(http.MethodDelete, srv.URL+"/api/recipes/"+pancakes.ID, nil)
	delResp, _ := http.DefaultClient.Do(delReq)
	delResp.Body.Close()

	// 7. Verify it's gone
	getResp, _ := http.Get(srv.URL + "/api/recipes/" + pancakes.ID)
	if getResp.StatusCode != http.StatusNotFound {
		t.Errorf("expected 404 after delete, got %d", getResp.StatusCode)
	}
	getResp.Body.Close()
}

func TestFullMealPlanWorkflow(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	// 1. Create recipes
	pancakes := createRecipe(t, srv, models.Recipe{
		Title:       "Pancakes",
		Ingredients: []models.Ingredient{{Name: "flour", Quantity: 2, Unit: "cups"}, {Name: "milk", Quantity: 1, Unit: "cup"}},
		Steps:       []string{"Mix", "Cook"},
		Servings:    4,
		Category:    "breakfast",
	})

	salad := createRecipe(t, srv, models.Recipe{
		Title:       "Caesar Salad",
		Ingredients: []models.Ingredient{{Name: "lettuce", Quantity: 1, Unit: "head"}, {Name: "croutons", Quantity: 1, Unit: "cup"}},
		Steps:       []string{"Chop", "Toss"},
		Servings:    2,
		Category:    "lunch",
	})

	steak := createRecipe(t, srv, models.Recipe{
		Title:       "Grilled Steak",
		Ingredients: []models.Ingredient{{Name: "beef steak", Quantity: 2, Unit: "pcs"}, {Name: "salt", Quantity: 1, Unit: "tsp"}},
		Steps:       []string{"Season", "Grill"},
		Servings:    2,
		Category:    "dinner",
	})

	// 2. Create a meal plan
	mpBody, _ := json.Marshal(models.MealPlan{Name: "Week 1"})
	mpResp, _ := http.Post(srv.URL+"/api/mealplans", "application/json", bytes.NewReader(mpBody))
	var mealPlan models.MealPlan
	json.NewDecoder(mpResp.Body).Decode(&mealPlan)
	mpResp.Body.Close()

	if mealPlan.ID == "" {
		t.Fatal("expected meal plan ID")
	}

	// 3. Add meals to the plan
	slots := []models.MealSlot{
		{Day: "monday", MealType: "breakfast", RecipeID: pancakes.ID},
		{Day: "monday", MealType: "lunch", RecipeID: salad.ID},
		{Day: "monday", MealType: "dinner", RecipeID: steak.ID},
		{Day: "tuesday", MealType: "breakfast", RecipeID: pancakes.ID},
	}

	for _, slot := range slots {
		slotBody, _ := json.Marshal(slot)
		slotResp, _ := http.Post(srv.URL+"/api/mealplans/"+mealPlan.ID+"/slots", "application/json", bytes.NewReader(slotBody))
		slotResp.Body.Close()
	}

	// 4. View the meal plan
	viewResp, _ := http.Get(srv.URL + "/api/mealplans/" + mealPlan.ID)
	var plan models.MealPlan
	json.NewDecoder(viewResp.Body).Decode(&plan)
	viewResp.Body.Close()

	if len(plan.Slots) != 4 {
		t.Errorf("expected 4 slots, got %d", len(plan.Slots))
	}

	// 5. Generate shopping list
	shopResp, _ := http.Get(srv.URL + "/api/mealplans/" + mealPlan.ID + "/shopping-list")
	var items []models.ShoppingItem
	json.NewDecoder(shopResp.Body).Decode(&items)
	shopResp.Body.Close()

	if len(items) == 0 {
		t.Error("expected shopping list items")
	}

	// Flour should be aggregated: 2 (monday breakfast) + 2 (tuesday breakfast) = 4
	flourFound := false
	for _, item := range items {
		if item.Name == "flour" {
			if item.Quantity != 4 {
				t.Errorf("expected flour quantity 4, got %f", item.Quantity)
			}
			flourFound = true
		}
	}
	if !flourFound {
		t.Error("expected flour in shopping list")
	}

	// 6. Remove a slot
	delReq, _ := http.NewRequest(http.MethodDelete, srv.URL+"/api/mealplans/"+mealPlan.ID+"/slots?day=tuesday&meal_type=breakfast", nil)
	delSlotResp, _ := http.DefaultClient.Do(delReq)
	delSlotResp.Body.Close()

	// 7. Verify slot removed
	viewResp2, _ := http.Get(srv.URL + "/api/mealplans/" + mealPlan.ID)
	var plan2 models.MealPlan
	json.NewDecoder(viewResp2.Body).Decode(&plan2)
	viewResp2.Body.Close()

	if len(plan2.Slots) != 3 {
		t.Errorf("expected 3 slots after removal, got %d", len(plan2.Slots))
	}

	// 8. Shopping list should now show flour = 2 (only monday breakfast)
	shopResp2, _ := http.Get(srv.URL + "/api/mealplans/" + mealPlan.ID + "/shopping-list")
	var items2 []models.ShoppingItem
	json.NewDecoder(shopResp2.Body).Decode(&items2)
	shopResp2.Body.Close()

	for _, item := range items2 {
		if item.Name == "flour" {
			if item.Quantity != 2 {
				t.Errorf("expected flour quantity 2 after slot removal, got %f", item.Quantity)
			}
		}
	}
}

func TestRecipeCategoryFiltering(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	// Create recipes in different categories
	createRecipe(t, srv, models.Recipe{
		Title: "Pancakes", Ingredients: []models.Ingredient{{Name: "flour", Quantity: 1, Unit: "cup"}},
		Steps: []string{"cook"}, Servings: 2, Category: "breakfast",
	})
	createRecipe(t, srv, models.Recipe{
		Title: "Steak", Ingredients: []models.Ingredient{{Name: "beef", Quantity: 1, Unit: "lb"}},
		Steps: []string{"grill"}, Servings: 2, Category: "dinner",
	})
	createRecipe(t, srv, models.Recipe{
		Title: "Brownie", Ingredients: []models.Ingredient{{Name: "chocolate", Quantity: 1, Unit: "bar"}},
		Steps: []string{"bake"}, Servings: 8, Category: "dessert",
	})

	// Filter by each category
	for _, tc := range []struct {
		category string
		expected int
	}{
		{"breakfast", 1},
		{"dinner", 1},
		{"dessert", 1},
		{"lunch", 0},
		{"snack", 0},
	} {
		resp, _ := http.Get(srv.URL + "/api/recipes?category=" + tc.category)
		var recipes []models.Recipe
		json.NewDecoder(resp.Body).Decode(&recipes)
		resp.Body.Close()
		if len(recipes) != tc.expected {
			t.Errorf("category %s: expected %d recipes, got %d", tc.category, tc.expected, len(recipes))
		}
	}
}
