package handlers

import (
	"net/http"
	"strings"

	"github.com/shootdaj/ax-test-recipes/pkg/models"
	"github.com/shootdaj/ax-test-recipes/pkg/store"
	"github.com/shootdaj/ax-test-recipes/pkg/utils"
)

// RecipeHandler handles recipe-related HTTP requests.
type RecipeHandler struct {
	Store *store.Store
}

// NewRecipeHandler creates a new RecipeHandler.
func NewRecipeHandler(s *store.Store) *RecipeHandler {
	return &RecipeHandler{Store: s}
}

// HandleRecipes handles GET (list) and POST (create) for /api/recipes.
func (h *RecipeHandler) HandleRecipes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.listRecipes(w, r)
	case http.MethodPost:
		h.createRecipe(w, r)
	default:
		utils.RespondError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// HandleRecipe handles GET, PUT, DELETE for /api/recipes/{id}.
func (h *RecipeHandler) HandleRecipe(w http.ResponseWriter, r *http.Request) {
	id := extractID(r.URL.Path, "/api/recipes/")
	if id == "" {
		utils.RespondError(w, http.StatusBadRequest, "recipe ID is required")
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getRecipe(w, id)
	case http.MethodPut:
		h.updateRecipe(w, r, id)
	case http.MethodDelete:
		h.deleteRecipe(w, id)
	default:
		utils.RespondError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *RecipeHandler) listRecipes(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	recipes := h.Store.ListRecipes(category)
	if recipes == nil {
		recipes = []*models.Recipe{}
	}
	utils.RespondJSON(w, http.StatusOK, recipes)
}

func (h *RecipeHandler) createRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	if err := utils.DecodeJSON(r, &recipe); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}

	if errs := recipe.Validate(); len(errs) > 0 {
		utils.RespondError(w, http.StatusBadRequest, strings.Join(errs, "; "))
		return
	}

	created := h.Store.CreateRecipe(&recipe)
	utils.RespondJSON(w, http.StatusCreated, created)
}

func (h *RecipeHandler) getRecipe(w http.ResponseWriter, id string) {
	recipe := h.Store.GetRecipe(id)
	if recipe == nil {
		utils.RespondError(w, http.StatusNotFound, "recipe not found")
		return
	}
	utils.RespondJSON(w, http.StatusOK, recipe)
}

func (h *RecipeHandler) updateRecipe(w http.ResponseWriter, r *http.Request, id string) {
	var recipe models.Recipe
	if err := utils.DecodeJSON(r, &recipe); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}

	if errs := recipe.Validate(); len(errs) > 0 {
		utils.RespondError(w, http.StatusBadRequest, strings.Join(errs, "; "))
		return
	}

	if !h.Store.UpdateRecipe(id, &recipe) {
		utils.RespondError(w, http.StatusNotFound, "recipe not found")
		return
	}
	recipe.ID = id
	utils.RespondJSON(w, http.StatusOK, &recipe)
}

func (h *RecipeHandler) deleteRecipe(w http.ResponseWriter, id string) {
	if !h.Store.DeleteRecipe(id) {
		utils.RespondError(w, http.StatusNotFound, "recipe not found")
		return
	}
	utils.RespondJSON(w, http.StatusNoContent, nil)
}

func extractID(path, prefix string) string {
	if !strings.HasPrefix(path, prefix) {
		return ""
	}
	id := strings.TrimPrefix(path, prefix)
	// Remove trailing slash and any further path segments
	if idx := strings.Index(id, "/"); idx != -1 {
		id = id[:idx]
	}
	return id
}
