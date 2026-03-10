package utils

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRespondJSON(t *testing.T) {
	w := httptest.NewRecorder()
	RespondJSON(w, http.StatusOK, map[string]string{"key": "value"})

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if ct := w.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", ct)
	}
	body := w.Body.String()
	if !strings.Contains(body, `"key":"value"`) {
		t.Errorf("unexpected body: %s", body)
	}
}

func TestRespondError(t *testing.T) {
	w := httptest.NewRecorder()
	RespondError(w, http.StatusBadRequest, "something went wrong")

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "something went wrong") {
		t.Errorf("expected error message in body: %s", body)
	}
}

func TestDecodeJSON(t *testing.T) {
	body := `{"name": "test"}`
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))

	var result struct {
		Name string `json:"name"`
	}
	err := DecodeJSON(r, &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Name != "test" {
		t.Errorf("expected 'test', got '%s'", result.Name)
	}
}

func TestDecodeJSON_InvalidJSON(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("not json"))

	var result struct{}
	err := DecodeJSON(r, &result)
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}
