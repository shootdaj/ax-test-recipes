package router

import (
	"net/http"
	"strings"

	"github.com/shootdaj/ax-test-recipes/internal/handlers"
	"github.com/shootdaj/ax-test-recipes/internal/store"
	"github.com/shootdaj/ax-test-recipes/pkg/utils"
)

// New creates a new HTTP handler with all routes configured.
func New(s *store.Store) http.Handler {
	recipeH := handlers.NewRecipeHandler(s)
	searchH := handlers.NewSearchHandler(s)
	mealPlanH := handlers.NewMealPlanHandler(s)

	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	// We need a catch-all handler for path-based routing
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// Enable CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Route matching
		switch {
		case path == "/api/recipes" || path == "/api/recipes/":
			recipeH.HandleRecipes(w, r)

		case path == "/api/search" || path == "/api/search/":
			searchH.HandleSearch(w, r)

		case strings.HasPrefix(path, "/api/recipes/") && strings.HasSuffix(path, "/scale"):
			searchH.HandleScale(w, r)

		case strings.HasPrefix(path, "/api/recipes/"):
			recipeH.HandleRecipe(w, r)

		case path == "/api/mealplans" || path == "/api/mealplans/":
			mealPlanH.HandleMealPlans(w, r)

		case strings.HasPrefix(path, "/api/mealplans/") && strings.HasSuffix(path, "/slots"):
			mealPlanH.HandleMealSlots(w, r)

		case strings.HasPrefix(path, "/api/mealplans/") && strings.HasSuffix(path, "/shopping-list"):
			mealPlanH.HandleShoppingList(w, r)

		case strings.HasPrefix(path, "/api/mealplans/"):
			mealPlanH.HandleMealPlan(w, r)

		case path == "/" || path == "":
			// Serve frontend (will be implemented in Phase 4)
			utils.RespondJSON(w, http.StatusOK, map[string]string{
				"name":    "Recipe Manager API",
				"version": "1.0.0",
				"status":  "ok",
			})

		default:
			utils.RespondError(w, http.StatusNotFound, "not found")
		}
	})

	return mux
}
