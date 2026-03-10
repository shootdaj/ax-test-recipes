package store

import (
	"fmt"
	"strings"
	"sync"

	"github.com/shootdaj/ax-test-recipes/pkg/models"
)

// Store provides thread-safe in-memory storage for recipes and meal plans.
type Store struct {
	mu        sync.RWMutex
	recipes   map[string]*models.Recipe
	mealPlans map[string]*models.MealPlan
	nextID    int
}

// New creates a new empty Store.
func New() *Store {
	return &Store{
		recipes:   make(map[string]*models.Recipe),
		mealPlans: make(map[string]*models.MealPlan),
		nextID:    1,
	}
}

func (s *Store) generateID() string {
	id := fmt.Sprintf("%d", s.nextID)
	s.nextID++
	return id
}

// CreateRecipe adds a new recipe and returns it with an assigned ID.
func (s *Store) CreateRecipe(r *models.Recipe) *models.Recipe {
	s.mu.Lock()
	defer s.mu.Unlock()

	r.ID = s.generateID()
	stored := *r
	s.recipes[r.ID] = &stored
	return &stored
}

// GetRecipe returns a recipe by ID, or nil if not found.
func (s *Store) GetRecipe(id string) *models.Recipe {
	s.mu.RLock()
	defer s.mu.RUnlock()

	r, ok := s.recipes[id]
	if !ok {
		return nil
	}
	copy := *r
	return &copy
}

// ListRecipes returns all recipes, optionally filtered by category.
func (s *Store) ListRecipes(category string) []*models.Recipe {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []*models.Recipe
	for _, r := range s.recipes {
		if category != "" && !strings.EqualFold(r.Category, category) {
			continue
		}
		copy := *r
		result = append(result, &copy)
	}
	return result
}

// UpdateRecipe updates an existing recipe. Returns false if not found.
func (s *Store) UpdateRecipe(id string, updated *models.Recipe) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.recipes[id]; !ok {
		return false
	}
	updated.ID = id
	stored := *updated
	s.recipes[id] = &stored
	return true
}

// DeleteRecipe removes a recipe by ID. Returns false if not found.
func (s *Store) DeleteRecipe(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.recipes[id]; !ok {
		return false
	}
	delete(s.recipes, id)
	return true
}

// SearchByIngredient returns recipes containing the given ingredient.
func (s *Store) SearchByIngredient(ingredient string) []*models.Recipe {
	s.mu.RLock()
	defer s.mu.RUnlock()

	ingredient = strings.ToLower(ingredient)
	var result []*models.Recipe
	for _, r := range s.recipes {
		for _, ing := range r.Ingredients {
			if strings.Contains(strings.ToLower(ing.Name), ingredient) {
				copy := *r
				result = append(result, &copy)
				break
			}
		}
	}
	return result
}

// SearchByKeyword returns recipes whose title or description contains the keyword.
func (s *Store) SearchByKeyword(keyword string) []*models.Recipe {
	s.mu.RLock()
	defer s.mu.RUnlock()

	keyword = strings.ToLower(keyword)
	var result []*models.Recipe
	for _, r := range s.recipes {
		if strings.Contains(strings.ToLower(r.Title), keyword) ||
			strings.Contains(strings.ToLower(r.Description), keyword) {
			copy := *r
			result = append(result, &copy)
		}
	}
	return result
}

// CreateMealPlan adds a new meal plan and returns it with an assigned ID.
func (s *Store) CreateMealPlan(mp *models.MealPlan) *models.MealPlan {
	s.mu.Lock()
	defer s.mu.Unlock()

	mp.ID = s.generateID()
	stored := *mp
	s.mealPlans[mp.ID] = &stored
	return &stored
}

// GetMealPlan returns a meal plan by ID.
func (s *Store) GetMealPlan(id string) *models.MealPlan {
	s.mu.RLock()
	defer s.mu.RUnlock()

	mp, ok := s.mealPlans[id]
	if !ok {
		return nil
	}
	copy := *mp
	return &copy
}

// ListMealPlans returns all meal plans.
func (s *Store) ListMealPlans() []*models.MealPlan {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []*models.MealPlan
	for _, mp := range s.mealPlans {
		copy := *mp
		result = append(result, &copy)
	}
	return result
}

// AddMealSlot adds a recipe to a meal plan slot. Returns false if meal plan not found.
func (s *Store) AddMealSlot(planID string, slot models.MealSlot) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	mp, ok := s.mealPlans[planID]
	if !ok {
		return false
	}

	// Remove existing slot for the same day+meal if present
	var filtered []models.MealSlot
	for _, existing := range mp.Slots {
		if !(strings.EqualFold(existing.Day, slot.Day) && strings.EqualFold(existing.MealType, slot.MealType)) {
			filtered = append(filtered, existing)
		}
	}
	mp.Slots = append(filtered, slot)
	return true
}

// RemoveMealSlot removes a recipe from a meal plan slot. Returns false if not found.
func (s *Store) RemoveMealSlot(planID, day, mealType string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	mp, ok := s.mealPlans[planID]
	if !ok {
		return false
	}

	var filtered []models.MealSlot
	found := false
	for _, existing := range mp.Slots {
		if strings.EqualFold(existing.Day, day) && strings.EqualFold(existing.MealType, mealType) {
			found = true
		} else {
			filtered = append(filtered, existing)
		}
	}
	mp.Slots = filtered
	return found
}

// GenerateShoppingList creates a shopping list from a meal plan by aggregating ingredients.
func (s *Store) GenerateShoppingList(planID string) ([]models.ShoppingItem, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	mp, ok := s.mealPlans[planID]
	if !ok {
		return nil, fmt.Errorf("meal plan not found")
	}

	// Aggregate ingredients by name+unit
	type key struct {
		name string
		unit string
	}
	aggregated := make(map[key]float64)
	order := []key{}

	for _, slot := range mp.Slots {
		recipe, ok := s.recipes[slot.RecipeID]
		if !ok {
			continue
		}
		for _, ing := range recipe.Ingredients {
			k := key{name: strings.ToLower(ing.Name), unit: strings.ToLower(ing.Unit)}
			if _, exists := aggregated[k]; !exists {
				order = append(order, k)
			}
			aggregated[k] += ing.Quantity
		}
	}

	var items []models.ShoppingItem
	for _, k := range order {
		items = append(items, models.ShoppingItem{
			Name:     k.name,
			Quantity: aggregated[k],
			Unit:     k.unit,
			Checked:  false,
		})
	}
	return items, nil
}
