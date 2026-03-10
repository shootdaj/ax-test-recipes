package models

// MealSlot represents a single meal assignment in a meal plan.
type MealSlot struct {
	Day      string `json:"day"`
	MealType string `json:"meal_type"`
	RecipeID string `json:"recipe_id"`
}

// MealPlan represents a weekly meal plan.
type MealPlan struct {
	ID    string     `json:"id"`
	Name  string     `json:"name"`
	Slots []MealSlot `json:"slots"`
}

// ShoppingItem represents an item on a shopping list.
type ShoppingItem struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
	Checked  bool    `json:"checked"`
}

// ValidDays defines the allowed days for meal planning.
var ValidDays = map[string]bool{
	"monday": true, "tuesday": true, "wednesday": true,
	"thursday": true, "friday": true, "saturday": true, "sunday": true,
}

// ValidMealTypes defines the allowed meal types.
var ValidMealTypes = map[string]bool{
	"breakfast": true, "lunch": true, "dinner": true, "snack": true,
}
