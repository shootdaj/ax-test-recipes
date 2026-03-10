package handlers

import (
	"net/http"
	"strconv"

	"github.com/shootdaj/ax-test-recipes/pkg/models"
	"github.com/shootdaj/ax-test-recipes/pkg/store"
	"github.com/shootdaj/ax-test-recipes/pkg/utils"
)

// SearchHandler handles search-related HTTP requests.
type SearchHandler struct {
	Store *store.Store
}

// NewSearchHandler creates a new SearchHandler.
func NewSearchHandler(s *store.Store) *SearchHandler {
	return &SearchHandler{Store: s}
}

// HandleSearch handles GET /api/search?ingredient=X or GET /api/search?q=keyword.
func (h *SearchHandler) HandleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	ingredient := r.URL.Query().Get("ingredient")
	keyword := r.URL.Query().Get("q")

	var recipes []*models.Recipe
	if ingredient != "" {
		recipes = h.Store.SearchByIngredient(ingredient)
	} else if keyword != "" {
		recipes = h.Store.SearchByKeyword(keyword)
	} else {
		utils.RespondError(w, http.StatusBadRequest, "provide 'ingredient' or 'q' query parameter")
		return
	}

	if recipes == nil {
		recipes = []*models.Recipe{}
	}
	utils.RespondJSON(w, http.StatusOK, recipes)
}

// HandleScale handles GET /api/recipes/{id}/scale?servings=N.
func (h *SearchHandler) HandleScale(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id := extractID(r.URL.Path, "/api/recipes/")
	// Remove /scale suffix from id
	if idx := len(id) - len("/scale"); idx > 0 && id[idx:] == "/scale" {
		id = id[:idx]
	}

	recipe := h.Store.GetRecipe(id)
	if recipe == nil {
		utils.RespondError(w, http.StatusNotFound, "recipe not found")
		return
	}

	servingsStr := r.URL.Query().Get("servings")
	if servingsStr == "" {
		utils.RespondError(w, http.StatusBadRequest, "servings query parameter is required")
		return
	}

	newServings, err := strconv.Atoi(servingsStr)
	if err != nil || newServings <= 0 {
		utils.RespondError(w, http.StatusBadRequest, "servings must be a positive integer")
		return
	}

	scaled := ScaleRecipe(recipe, newServings)
	utils.RespondJSON(w, http.StatusOK, scaled)
}

// ScaleRecipe returns a copy of the recipe with ingredient quantities adjusted for new servings.
func ScaleRecipe(recipe *models.Recipe, newServings int) *models.Recipe {
	if recipe.Servings == 0 {
		return recipe
	}

	factor := float64(newServings) / float64(recipe.Servings)
	scaled := *recipe
	scaled.Servings = newServings
	scaled.Ingredients = make([]models.Ingredient, len(recipe.Ingredients))
	for i, ing := range recipe.Ingredients {
		scaled.Ingredients[i] = models.Ingredient{
			Name:     ing.Name,
			Quantity: ing.Quantity * factor,
			Unit:     ing.Unit,
		}
	}
	return &scaled
}
