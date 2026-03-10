package handler

import (
	"net/http"

	"github.com/shootdaj/ax-test-recipes/pkg/router"
	"github.com/shootdaj/ax-test-recipes/pkg/store"
)

var (
	appStore   *store.Store
	appHandler http.Handler
)

func init() {
	appStore = store.New()
	appHandler = router.New(appStore)
}

// Handler is the Vercel serverless function entry point.
func Handler(w http.ResponseWriter, r *http.Request) {
	appHandler.ServeHTTP(w, r)
}
