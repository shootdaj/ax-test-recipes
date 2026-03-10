//go:build integration

package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shootdaj/ax-test-recipes/pkg/models"
	"github.com/shootdaj/ax-test-recipes/pkg/router"
	"github.com/shootdaj/ax-test-recipes/pkg/store"
)

func setupServer() *httptest.Server {
	s := store.New()
	handler := router.New(s)
	return httptest.NewServer(handler)
}

func TestAPI_HealthEndpoint(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/health")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestAPI_CreateAndGetRecipe(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	recipe := models.Recipe{
		Title:       "Pancakes",
		Description: "Fluffy pancakes",
		Ingredients: []models.Ingredient{{Name: "flour", Quantity: 2, Unit: "cups"}},
		Steps:       []string{"Mix", "Cook"},
		Servings:    4,
		Category:    "breakfast",
	}

	body, _ := json.Marshal(recipe)
	resp, err := http.Post(srv.URL+"/api/recipes", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected 201, got %d", resp.StatusCode)
	}

	var created models.Recipe
	json.NewDecoder(resp.Body).Decode(&created)
	if created.ID == "" {
		t.Error("expected ID in response")
	}

	// Get the recipe
	getResp, err := http.Get(srv.URL + "/api/recipes/" + created.ID)
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}
	defer getResp.Body.Close()

	if getResp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", getResp.StatusCode)
	}

	var got models.Recipe
	json.NewDecoder(getResp.Body).Decode(&got)
	if got.Title != "Pancakes" {
		t.Errorf("expected 'Pancakes', got '%s'", got.Title)
	}
}

func TestAPI_ListRecipes_FilterByCategory(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	// Create two recipes
	for _, r := range []models.Recipe{
		{Title: "Pancakes", Ingredients: []models.Ingredient{{Name: "flour", Quantity: 1, Unit: "cup"}}, Steps: []string{"cook"}, Servings: 2, Category: "breakfast"},
		{Title: "Steak", Ingredients: []models.Ingredient{{Name: "beef", Quantity: 1, Unit: "lb"}}, Steps: []string{"grill"}, Servings: 2, Category: "dinner"},
	} {
		body, _ := json.Marshal(r)
		http.Post(srv.URL+"/api/recipes", "application/json", bytes.NewReader(body))
	}

	// Filter by breakfast
	resp, _ := http.Get(srv.URL + "/api/recipes?category=breakfast")
	defer resp.Body.Close()

	var recipes []models.Recipe
	json.NewDecoder(resp.Body).Decode(&recipes)
	if len(recipes) != 1 {
		t.Errorf("expected 1 breakfast recipe, got %d", len(recipes))
	}
}

func TestAPI_UpdateRecipe(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	// Create
	recipe := models.Recipe{
		Title: "Pancakes", Ingredients: []models.Ingredient{{Name: "flour", Quantity: 1, Unit: "cup"}},
		Steps: []string{"cook"}, Servings: 2, Category: "breakfast",
	}
	body, _ := json.Marshal(recipe)
	resp, _ := http.Post(srv.URL+"/api/recipes", "application/json", bytes.NewReader(body))
	var created models.Recipe
	json.NewDecoder(resp.Body).Decode(&created)
	resp.Body.Close()

	// Update
	recipe.Title = "Super Pancakes"
	body, _ = json.Marshal(recipe)
	req, _ := http.NewRequest(http.MethodPut, srv.URL+"/api/recipes/"+created.ID, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	updateResp, _ := http.DefaultClient.Do(req)
	defer updateResp.Body.Close()

	if updateResp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", updateResp.StatusCode)
	}

	var updated models.Recipe
	json.NewDecoder(updateResp.Body).Decode(&updated)
	if updated.Title != "Super Pancakes" {
		t.Errorf("expected 'Super Pancakes', got '%s'", updated.Title)
	}
}

func TestAPI_DeleteRecipe(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	// Create
	recipe := models.Recipe{
		Title: "Temp", Ingredients: []models.Ingredient{{Name: "x", Quantity: 1, Unit: "u"}},
		Steps: []string{"do"}, Servings: 1,
	}
	body, _ := json.Marshal(recipe)
	resp, _ := http.Post(srv.URL+"/api/recipes", "application/json", bytes.NewReader(body))
	var created models.Recipe
	json.NewDecoder(resp.Body).Decode(&created)
	resp.Body.Close()

	// Delete
	req, _ := http.NewRequest(http.MethodDelete, srv.URL+"/api/recipes/"+created.ID, nil)
	delResp, _ := http.DefaultClient.Do(req)
	defer delResp.Body.Close()

	if delResp.StatusCode != http.StatusNoContent {
		t.Errorf("expected 204, got %d", delResp.StatusCode)
	}

	// Verify gone
	getResp, _ := http.Get(srv.URL + "/api/recipes/" + created.ID)
	defer getResp.Body.Close()
	if getResp.StatusCode != http.StatusNotFound {
		t.Errorf("expected 404 after delete, got %d", getResp.StatusCode)
	}
}

func TestAPI_SearchByIngredient(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	recipe := models.Recipe{
		Title: "Pancakes", Ingredients: []models.Ingredient{{Name: "flour", Quantity: 1, Unit: "cup"}},
		Steps: []string{"cook"}, Servings: 2,
	}
	body, _ := json.Marshal(recipe)
	http.Post(srv.URL+"/api/recipes", "application/json", bytes.NewReader(body))

	resp, _ := http.Get(srv.URL + "/api/search?ingredient=flour")
	defer resp.Body.Close()

	var results []models.Recipe
	json.NewDecoder(resp.Body).Decode(&results)
	if len(results) != 1 {
		t.Errorf("expected 1 result, got %d", len(results))
	}
}

func TestAPI_SearchByKeyword(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	recipe := models.Recipe{
		Title: "Chocolate Cake", Ingredients: []models.Ingredient{{Name: "cocoa", Quantity: 1, Unit: "cup"}},
		Steps: []string{"bake"}, Servings: 8,
	}
	body, _ := json.Marshal(recipe)
	http.Post(srv.URL+"/api/recipes", "application/json", bytes.NewReader(body))

	resp, _ := http.Get(srv.URL + "/api/search?q=chocolate")
	defer resp.Body.Close()

	var results []models.Recipe
	json.NewDecoder(resp.Body).Decode(&results)
	if len(results) != 1 {
		t.Errorf("expected 1 result, got %d", len(results))
	}
}

func TestAPI_ScaleRecipe(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	recipe := models.Recipe{
		Title: "Pancakes", Ingredients: []models.Ingredient{{Name: "flour", Quantity: 2, Unit: "cups"}},
		Steps: []string{"cook"}, Servings: 4,
	}
	body, _ := json.Marshal(recipe)
	resp, _ := http.Post(srv.URL+"/api/recipes", "application/json", bytes.NewReader(body))
	var created models.Recipe
	json.NewDecoder(resp.Body).Decode(&created)
	resp.Body.Close()

	scaleResp, _ := http.Get(srv.URL + "/api/recipes/" + created.ID + "/scale?servings=8")
	defer scaleResp.Body.Close()

	if scaleResp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", scaleResp.StatusCode)
	}

	var scaled models.Recipe
	json.NewDecoder(scaleResp.Body).Decode(&scaled)
	if scaled.Servings != 8 {
		t.Errorf("expected 8 servings, got %d", scaled.Servings)
	}
	if scaled.Ingredients[0].Quantity != 4 {
		t.Errorf("expected flour 4, got %f", scaled.Ingredients[0].Quantity)
	}
}
