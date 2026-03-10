package handlers

import (
	"net/http"
	"strings"

	"github.com/shootdaj/ax-test-recipes/pkg/models"
	"github.com/shootdaj/ax-test-recipes/pkg/store"
	"github.com/shootdaj/ax-test-recipes/pkg/utils"
)

// MealPlanHandler handles meal plan HTTP requests.
type MealPlanHandler struct {
	Store *store.Store
}

// NewMealPlanHandler creates a new MealPlanHandler.
func NewMealPlanHandler(s *store.Store) *MealPlanHandler {
	return &MealPlanHandler{Store: s}
}

// HandleMealPlans handles GET (list) and POST (create) for /api/mealplans.
func (h *MealPlanHandler) HandleMealPlans(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.listMealPlans(w)
	case http.MethodPost:
		h.createMealPlan(w, r)
	default:
		utils.RespondError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// HandleMealPlan handles GET for /api/mealplans/{id}.
func (h *MealPlanHandler) HandleMealPlan(w http.ResponseWriter, r *http.Request) {
	id := extractID(r.URL.Path, "/api/mealplans/")
	if id == "" {
		utils.RespondError(w, http.StatusBadRequest, "meal plan ID is required")
		return
	}

	if r.Method != http.MethodGet {
		utils.RespondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	mp := h.Store.GetMealPlan(id)
	if mp == nil {
		utils.RespondError(w, http.StatusNotFound, "meal plan not found")
		return
	}
	utils.RespondJSON(w, http.StatusOK, mp)
}

// HandleMealSlots handles POST (add) and DELETE (remove) for /api/mealplans/{id}/slots.
func (h *MealPlanHandler) HandleMealSlots(w http.ResponseWriter, r *http.Request) {
	// Extract plan ID from path like /api/mealplans/{id}/slots
	path := strings.TrimPrefix(r.URL.Path, "/api/mealplans/")
	parts := strings.SplitN(path, "/", 2)
	if len(parts) < 1 || parts[0] == "" {
		utils.RespondError(w, http.StatusBadRequest, "meal plan ID is required")
		return
	}
	planID := parts[0]

	switch r.Method {
	case http.MethodPost:
		h.addSlot(w, r, planID)
	case http.MethodDelete:
		h.removeSlot(w, r, planID)
	default:
		utils.RespondError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// HandleShoppingList handles GET /api/mealplans/{id}/shopping-list.
func (h *MealPlanHandler) HandleShoppingList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Extract plan ID from path like /api/mealplans/{id}/shopping-list
	path := strings.TrimPrefix(r.URL.Path, "/api/mealplans/")
	parts := strings.SplitN(path, "/", 2)
	if len(parts) < 1 || parts[0] == "" {
		utils.RespondError(w, http.StatusBadRequest, "meal plan ID is required")
		return
	}
	planID := parts[0]

	items, err := h.Store.GenerateShoppingList(planID)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	if items == nil {
		items = []models.ShoppingItem{}
	}
	utils.RespondJSON(w, http.StatusOK, items)
}

func (h *MealPlanHandler) listMealPlans(w http.ResponseWriter) {
	plans := h.Store.ListMealPlans()
	if plans == nil {
		plans = []*models.MealPlan{}
	}
	utils.RespondJSON(w, http.StatusOK, plans)
}

func (h *MealPlanHandler) createMealPlan(w http.ResponseWriter, r *http.Request) {
	var mp models.MealPlan
	if err := utils.DecodeJSON(r, &mp); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	if mp.Name == "" {
		utils.RespondError(w, http.StatusBadRequest, "name is required")
		return
	}
	created := h.Store.CreateMealPlan(&mp)
	utils.RespondJSON(w, http.StatusCreated, created)
}

func (h *MealPlanHandler) addSlot(w http.ResponseWriter, r *http.Request, planID string) {
	var slot models.MealSlot
	if err := utils.DecodeJSON(r, &slot); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}

	if !models.ValidDays[strings.ToLower(slot.Day)] {
		utils.RespondError(w, http.StatusBadRequest, "invalid day")
		return
	}
	if !models.ValidMealTypes[strings.ToLower(slot.MealType)] {
		utils.RespondError(w, http.StatusBadRequest, "invalid meal type")
		return
	}
	if slot.RecipeID == "" {
		utils.RespondError(w, http.StatusBadRequest, "recipe_id is required")
		return
	}

	// Verify recipe exists
	if h.Store.GetRecipe(slot.RecipeID) == nil {
		utils.RespondError(w, http.StatusBadRequest, "recipe not found")
		return
	}

	if !h.Store.AddMealSlot(planID, slot) {
		utils.RespondError(w, http.StatusNotFound, "meal plan not found")
		return
	}

	mp := h.Store.GetMealPlan(planID)
	utils.RespondJSON(w, http.StatusOK, mp)
}

func (h *MealPlanHandler) removeSlot(w http.ResponseWriter, r *http.Request, planID string) {
	day := r.URL.Query().Get("day")
	mealType := r.URL.Query().Get("meal_type")

	if day == "" || mealType == "" {
		utils.RespondError(w, http.StatusBadRequest, "day and meal_type query parameters are required")
		return
	}

	if !h.Store.RemoveMealSlot(planID, day, mealType) {
		utils.RespondError(w, http.StatusNotFound, "meal slot not found")
		return
	}

	mp := h.Store.GetMealPlan(planID)
	utils.RespondJSON(w, http.StatusOK, mp)
}
