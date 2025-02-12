package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestItemsHandlerGet(t *testing.T) {
	// Reset global items to an empty slice for testing.
	items = []Item{}

	// Create a new HTTP GET request.
	req := httptest.NewRequest(http.MethodGet, "/items", nil)
	w := httptest.NewRecorder()

	// Call the itemsHandler.
	itemsHandler(w, req)
	res := w.Result()
	defer res.Body.Close()

	// Check for the expected status code.
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.StatusCode)
	}

	// Decode the response body.
	var responseItems []Item
	err := json.NewDecoder(res.Body).Decode(&responseItems)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	// Since we reset items to empty, we expect an empty slice.
	if len(responseItems) != 0 {
		t.Fatalf("expected no items, got %d", len(responseItems))
	}
}
