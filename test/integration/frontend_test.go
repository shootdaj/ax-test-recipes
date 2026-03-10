//go:build integration

package integration

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestAPI_FrontendServesHTML(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}

	ct := resp.Header.Get("Content-Type")
	if !strings.Contains(ct, "text/html") {
		t.Errorf("expected text/html content type, got %s", ct)
	}

	body, _ := io.ReadAll(resp.Body)
	bodyStr := string(body)

	if !strings.Contains(bodyStr, "Recipe Manager") {
		t.Error("expected 'Recipe Manager' in HTML")
	}
	if !strings.Contains(bodyStr, "page-recipes") {
		t.Error("expected recipe page elements in HTML")
	}
	if !strings.Contains(bodyStr, "page-mealplan") {
		t.Error("expected meal plan page elements in HTML")
	}
	if !strings.Contains(bodyStr, "page-shopping") {
		t.Error("expected shopping list page elements in HTML")
	}
}
