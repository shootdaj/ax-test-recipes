package models

// Ingredient represents a single ingredient in a recipe.
type Ingredient struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

// Recipe represents a cooking recipe with all its details.
type Recipe struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Ingredients []Ingredient `json:"ingredients"`
	Steps       []string     `json:"steps"`
	PrepTime    int          `json:"prep_time"`
	CookTime    int          `json:"cook_time"`
	Servings    int          `json:"servings"`
	Tags        []string     `json:"tags"`
	ImageURL    string       `json:"image_url"`
	Category    string       `json:"category"`
}

// ValidCategories defines the allowed recipe categories.
var ValidCategories = map[string]bool{
	"breakfast": true,
	"lunch":     true,
	"dinner":    true,
	"dessert":   true,
	"snack":     true,
}

// Validate checks that a recipe has the required fields.
func (r *Recipe) Validate() []string {
	var errors []string
	if r.Title == "" {
		errors = append(errors, "title is required")
	}
	if len(r.Ingredients) == 0 {
		errors = append(errors, "at least one ingredient is required")
	}
	if len(r.Steps) == 0 {
		errors = append(errors, "at least one step is required")
	}
	if r.Servings <= 0 {
		errors = append(errors, "servings must be positive")
	}
	if r.Category != "" && !ValidCategories[r.Category] {
		errors = append(errors, "category must be one of: breakfast, lunch, dinner, dessert, snack")
	}
	return errors
}
