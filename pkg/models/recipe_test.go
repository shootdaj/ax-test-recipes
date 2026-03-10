package models

import "testing"

func TestRecipe_Validate_ValidRecipe(t *testing.T) {
	r := &Recipe{
		Title:       "Pancakes",
		Ingredients: []Ingredient{{Name: "flour", Quantity: 2, Unit: "cups"}},
		Steps:       []string{"Mix", "Cook"},
		Servings:    4,
		Category:    "breakfast",
	}
	errs := r.Validate()
	if len(errs) != 0 {
		t.Errorf("expected no errors, got: %v", errs)
	}
}

func TestRecipe_Validate_MissingTitle(t *testing.T) {
	r := &Recipe{
		Ingredients: []Ingredient{{Name: "flour", Quantity: 2, Unit: "cups"}},
		Steps:       []string{"Mix"},
		Servings:    4,
	}
	errs := r.Validate()
	if len(errs) == 0 {
		t.Error("expected validation error for missing title")
	}
	found := false
	for _, e := range errs {
		if e == "title is required" {
			found = true
		}
	}
	if !found {
		t.Error("expected 'title is required' error")
	}
}

func TestRecipe_Validate_MissingIngredients(t *testing.T) {
	r := &Recipe{
		Title:    "Pancakes",
		Steps:    []string{"Mix"},
		Servings: 4,
	}
	errs := r.Validate()
	found := false
	for _, e := range errs {
		if e == "at least one ingredient is required" {
			found = true
		}
	}
	if !found {
		t.Error("expected ingredient validation error")
	}
}

func TestRecipe_Validate_MissingSteps(t *testing.T) {
	r := &Recipe{
		Title:       "Pancakes",
		Ingredients: []Ingredient{{Name: "flour", Quantity: 2, Unit: "cups"}},
		Servings:    4,
	}
	errs := r.Validate()
	found := false
	for _, e := range errs {
		if e == "at least one step is required" {
			found = true
		}
	}
	if !found {
		t.Error("expected step validation error")
	}
}

func TestRecipe_Validate_InvalidServings(t *testing.T) {
	r := &Recipe{
		Title:       "Pancakes",
		Ingredients: []Ingredient{{Name: "flour", Quantity: 2, Unit: "cups"}},
		Steps:       []string{"Mix"},
		Servings:    0,
	}
	errs := r.Validate()
	found := false
	for _, e := range errs {
		if e == "servings must be positive" {
			found = true
		}
	}
	if !found {
		t.Error("expected servings validation error")
	}
}

func TestRecipe_Validate_InvalidCategory(t *testing.T) {
	r := &Recipe{
		Title:       "Pancakes",
		Ingredients: []Ingredient{{Name: "flour", Quantity: 2, Unit: "cups"}},
		Steps:       []string{"Mix"},
		Servings:    4,
		Category:    "invalid",
	}
	errs := r.Validate()
	found := false
	for _, e := range errs {
		if e == "category must be one of: breakfast, lunch, dinner, dessert, snack" {
			found = true
		}
	}
	if !found {
		t.Error("expected category validation error")
	}
}

func TestRecipe_Validate_EmptyCategory(t *testing.T) {
	r := &Recipe{
		Title:       "Pancakes",
		Ingredients: []Ingredient{{Name: "flour", Quantity: 2, Unit: "cups"}},
		Steps:       []string{"Mix"},
		Servings:    4,
		Category:    "",
	}
	errs := r.Validate()
	if len(errs) != 0 {
		t.Errorf("empty category should be valid, got: %v", errs)
	}
}

func TestRecipe_Validate_MultipleErrors(t *testing.T) {
	r := &Recipe{}
	errs := r.Validate()
	if len(errs) < 3 {
		t.Errorf("expected at least 3 errors, got %d: %v", len(errs), errs)
	}
}
