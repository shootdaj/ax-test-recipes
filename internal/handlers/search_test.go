package handlers

import (
	"testing"

	"github.com/shootdaj/ax-test-recipes/internal/models"
)

func TestScaleRecipe_DoubleServings(t *testing.T) {
	recipe := &models.Recipe{
		Title:    "Pancakes",
		Servings: 4,
		Ingredients: []models.Ingredient{
			{Name: "flour", Quantity: 2, Unit: "cups"},
			{Name: "sugar", Quantity: 1, Unit: "cup"},
		},
	}

	scaled := ScaleRecipe(recipe, 8)
	if scaled.Servings != 8 {
		t.Errorf("expected 8 servings, got %d", scaled.Servings)
	}
	if scaled.Ingredients[0].Quantity != 4 {
		t.Errorf("expected flour quantity 4, got %f", scaled.Ingredients[0].Quantity)
	}
	if scaled.Ingredients[1].Quantity != 2 {
		t.Errorf("expected sugar quantity 2, got %f", scaled.Ingredients[1].Quantity)
	}
}

func TestScaleRecipe_HalfServings(t *testing.T) {
	recipe := &models.Recipe{
		Title:    "Cake",
		Servings: 8,
		Ingredients: []models.Ingredient{
			{Name: "flour", Quantity: 4, Unit: "cups"},
		},
	}

	scaled := ScaleRecipe(recipe, 4)
	if scaled.Servings != 4 {
		t.Errorf("expected 4 servings, got %d", scaled.Servings)
	}
	if scaled.Ingredients[0].Quantity != 2 {
		t.Errorf("expected flour quantity 2, got %f", scaled.Ingredients[0].Quantity)
	}
}

func TestScaleRecipe_SameServings(t *testing.T) {
	recipe := &models.Recipe{
		Title:    "Pancakes",
		Servings: 4,
		Ingredients: []models.Ingredient{
			{Name: "flour", Quantity: 2, Unit: "cups"},
		},
	}

	scaled := ScaleRecipe(recipe, 4)
	if scaled.Ingredients[0].Quantity != 2 {
		t.Errorf("expected same quantity, got %f", scaled.Ingredients[0].Quantity)
	}
}

func TestScaleRecipe_ZeroOriginalServings(t *testing.T) {
	recipe := &models.Recipe{
		Title:    "Test",
		Servings: 0,
		Ingredients: []models.Ingredient{
			{Name: "flour", Quantity: 2, Unit: "cups"},
		},
	}

	scaled := ScaleRecipe(recipe, 4)
	// Should return unchanged when original servings is 0
	if scaled.Ingredients[0].Quantity != 2 {
		t.Errorf("expected unchanged quantity for 0 original servings")
	}
}

func TestScaleRecipe_DoesNotMutateOriginal(t *testing.T) {
	recipe := &models.Recipe{
		Title:    "Pancakes",
		Servings: 4,
		Ingredients: []models.Ingredient{
			{Name: "flour", Quantity: 2, Unit: "cups"},
		},
	}

	ScaleRecipe(recipe, 8)
	if recipe.Ingredients[0].Quantity != 2 {
		t.Error("original recipe was mutated by scaling")
	}
	if recipe.Servings != 4 {
		t.Error("original servings was mutated by scaling")
	}
}
