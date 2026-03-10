package store

import (
	"testing"

	"github.com/shootdaj/ax-test-recipes/internal/models"
)

func newTestRecipe(title, category string) *models.Recipe {
	return &models.Recipe{
		Title:       title,
		Description: "Test recipe",
		Ingredients: []models.Ingredient{
			{Name: "flour", Quantity: 2, Unit: "cups"},
			{Name: "sugar", Quantity: 1, Unit: "cup"},
		},
		Steps:    []string{"Mix", "Cook"},
		PrepTime: 10,
		CookTime: 20,
		Servings: 4,
		Tags:     []string{"test"},
		Category: category,
	}
}

func TestStore_CreateRecipe(t *testing.T) {
	s := New()
	r := newTestRecipe("Pancakes", "breakfast")
	created := s.CreateRecipe(r)

	if created.ID == "" {
		t.Error("expected ID to be assigned")
	}
	if created.Title != "Pancakes" {
		t.Errorf("expected title 'Pancakes', got '%s'", created.Title)
	}
}

func TestStore_GetRecipe(t *testing.T) {
	s := New()
	created := s.CreateRecipe(newTestRecipe("Pancakes", "breakfast"))

	got := s.GetRecipe(created.ID)
	if got == nil {
		t.Fatal("expected recipe, got nil")
	}
	if got.Title != "Pancakes" {
		t.Errorf("expected 'Pancakes', got '%s'", got.Title)
	}
}

func TestStore_GetRecipe_NotFound(t *testing.T) {
	s := New()
	got := s.GetRecipe("nonexistent")
	if got != nil {
		t.Error("expected nil for nonexistent recipe")
	}
}

func TestStore_ListRecipes(t *testing.T) {
	s := New()
	s.CreateRecipe(newTestRecipe("Pancakes", "breakfast"))
	s.CreateRecipe(newTestRecipe("Steak", "dinner"))
	s.CreateRecipe(newTestRecipe("Eggs", "breakfast"))

	all := s.ListRecipes("")
	if len(all) != 3 {
		t.Errorf("expected 3 recipes, got %d", len(all))
	}
}

func TestStore_ListRecipes_FilterByCategory(t *testing.T) {
	s := New()
	s.CreateRecipe(newTestRecipe("Pancakes", "breakfast"))
	s.CreateRecipe(newTestRecipe("Steak", "dinner"))
	s.CreateRecipe(newTestRecipe("Eggs", "breakfast"))

	breakfast := s.ListRecipes("breakfast")
	if len(breakfast) != 2 {
		t.Errorf("expected 2 breakfast recipes, got %d", len(breakfast))
	}
}

func TestStore_UpdateRecipe(t *testing.T) {
	s := New()
	created := s.CreateRecipe(newTestRecipe("Pancakes", "breakfast"))

	updated := newTestRecipe("Super Pancakes", "breakfast")
	ok := s.UpdateRecipe(created.ID, updated)
	if !ok {
		t.Error("expected update to succeed")
	}

	got := s.GetRecipe(created.ID)
	if got.Title != "Super Pancakes" {
		t.Errorf("expected 'Super Pancakes', got '%s'", got.Title)
	}
}

func TestStore_UpdateRecipe_NotFound(t *testing.T) {
	s := New()
	ok := s.UpdateRecipe("nonexistent", newTestRecipe("X", "dinner"))
	if ok {
		t.Error("expected update to fail for nonexistent recipe")
	}
}

func TestStore_DeleteRecipe(t *testing.T) {
	s := New()
	created := s.CreateRecipe(newTestRecipe("Pancakes", "breakfast"))

	ok := s.DeleteRecipe(created.ID)
	if !ok {
		t.Error("expected delete to succeed")
	}

	got := s.GetRecipe(created.ID)
	if got != nil {
		t.Error("expected recipe to be deleted")
	}
}

func TestStore_DeleteRecipe_NotFound(t *testing.T) {
	s := New()
	ok := s.DeleteRecipe("nonexistent")
	if ok {
		t.Error("expected delete to fail for nonexistent recipe")
	}
}

func TestStore_SearchByIngredient(t *testing.T) {
	s := New()
	s.CreateRecipe(newTestRecipe("Pancakes", "breakfast")) // has flour, sugar
	cake := newTestRecipe("Cake", "dessert")
	cake.Ingredients = []models.Ingredient{{Name: "chocolate", Quantity: 1, Unit: "bar"}}
	s.CreateRecipe(cake)

	results := s.SearchByIngredient("flour")
	if len(results) != 1 {
		t.Errorf("expected 1 result for 'flour', got %d", len(results))
	}
	if results[0].Title != "Pancakes" {
		t.Errorf("expected 'Pancakes', got '%s'", results[0].Title)
	}
}

func TestStore_SearchByIngredient_CaseInsensitive(t *testing.T) {
	s := New()
	s.CreateRecipe(newTestRecipe("Pancakes", "breakfast"))

	results := s.SearchByIngredient("FLOUR")
	if len(results) != 1 {
		t.Errorf("expected 1 result for 'FLOUR', got %d", len(results))
	}
}

func TestStore_SearchByKeyword(t *testing.T) {
	s := New()
	s.CreateRecipe(newTestRecipe("Pancakes", "breakfast"))
	s.CreateRecipe(newTestRecipe("Steak", "dinner"))

	results := s.SearchByKeyword("pancake")
	if len(results) != 1 {
		t.Errorf("expected 1 result for 'pancake', got %d", len(results))
	}
}

func TestStore_CreateMealPlan(t *testing.T) {
	s := New()
	mp := &models.MealPlan{Name: "Week 1"}
	created := s.CreateMealPlan(mp)

	if created.ID == "" {
		t.Error("expected ID to be assigned")
	}
	if created.Name != "Week 1" {
		t.Errorf("expected 'Week 1', got '%s'", created.Name)
	}
}

func TestStore_AddMealSlot(t *testing.T) {
	s := New()
	recipe := s.CreateRecipe(newTestRecipe("Pancakes", "breakfast"))
	mp := s.CreateMealPlan(&models.MealPlan{Name: "Week 1"})

	ok := s.AddMealSlot(mp.ID, models.MealSlot{
		Day:      "monday",
		MealType: "breakfast",
		RecipeID: recipe.ID,
	})
	if !ok {
		t.Error("expected add slot to succeed")
	}

	got := s.GetMealPlan(mp.ID)
	if len(got.Slots) != 1 {
		t.Errorf("expected 1 slot, got %d", len(got.Slots))
	}
}

func TestStore_AddMealSlot_ReplacesExisting(t *testing.T) {
	s := New()
	r1 := s.CreateRecipe(newTestRecipe("Pancakes", "breakfast"))
	r2 := s.CreateRecipe(newTestRecipe("Eggs", "breakfast"))
	mp := s.CreateMealPlan(&models.MealPlan{Name: "Week 1"})

	s.AddMealSlot(mp.ID, models.MealSlot{Day: "monday", MealType: "breakfast", RecipeID: r1.ID})
	s.AddMealSlot(mp.ID, models.MealSlot{Day: "monday", MealType: "breakfast", RecipeID: r2.ID})

	got := s.GetMealPlan(mp.ID)
	if len(got.Slots) != 1 {
		t.Errorf("expected 1 slot (replaced), got %d", len(got.Slots))
	}
	if got.Slots[0].RecipeID != r2.ID {
		t.Error("expected slot to have the second recipe")
	}
}

func TestStore_RemoveMealSlot(t *testing.T) {
	s := New()
	recipe := s.CreateRecipe(newTestRecipe("Pancakes", "breakfast"))
	mp := s.CreateMealPlan(&models.MealPlan{Name: "Week 1"})
	s.AddMealSlot(mp.ID, models.MealSlot{Day: "monday", MealType: "breakfast", RecipeID: recipe.ID})

	ok := s.RemoveMealSlot(mp.ID, "monday", "breakfast")
	if !ok {
		t.Error("expected remove to succeed")
	}

	got := s.GetMealPlan(mp.ID)
	if len(got.Slots) != 0 {
		t.Errorf("expected 0 slots, got %d", len(got.Slots))
	}
}

func TestStore_GenerateShoppingList(t *testing.T) {
	s := New()
	r1 := s.CreateRecipe(newTestRecipe("Pancakes", "breakfast")) // flour 2, sugar 1
	r2 := newTestRecipe("Cake", "dessert")
	r2.Ingredients = []models.Ingredient{
		{Name: "flour", Quantity: 3, Unit: "cups"},
		{Name: "butter", Quantity: 0.5, Unit: "cups"},
	}
	r2Created := s.CreateRecipe(r2)

	mp := s.CreateMealPlan(&models.MealPlan{Name: "Week 1"})
	s.AddMealSlot(mp.ID, models.MealSlot{Day: "monday", MealType: "breakfast", RecipeID: r1.ID})
	s.AddMealSlot(mp.ID, models.MealSlot{Day: "tuesday", MealType: "dessert", RecipeID: r2Created.ID})

	items, err := s.GenerateShoppingList(mp.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// flour should be aggregated: 2 + 3 = 5
	flourFound := false
	for _, item := range items {
		if item.Name == "flour" && item.Unit == "cups" {
			if item.Quantity != 5 {
				t.Errorf("expected flour quantity 5, got %f", item.Quantity)
			}
			flourFound = true
		}
	}
	if !flourFound {
		t.Error("expected flour in shopping list")
	}

	// sugar should be 1
	sugarFound := false
	for _, item := range items {
		if item.Name == "sugar" && item.Unit == "cup" {
			if item.Quantity != 1 {
				t.Errorf("expected sugar quantity 1, got %f", item.Quantity)
			}
			sugarFound = true
		}
	}
	if !sugarFound {
		t.Error("expected sugar in shopping list")
	}

	// butter should be 0.5
	butterFound := false
	for _, item := range items {
		if item.Name == "butter" && item.Unit == "cups" {
			if item.Quantity != 0.5 {
				t.Errorf("expected butter quantity 0.5, got %f", item.Quantity)
			}
			butterFound = true
		}
	}
	if !butterFound {
		t.Error("expected butter in shopping list")
	}
}

func TestStore_GenerateShoppingList_NotFound(t *testing.T) {
	s := New()
	_, err := s.GenerateShoppingList("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent meal plan")
	}
}
