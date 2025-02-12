package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Global in-memory storage for items and a simple counter for IDs.
var items = []Item{}
var nextID int = 1

// itemsHandler handles requests to /items.
func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Return the list of items.
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	case http.MethodPost:
		// Create a new item.
		var newItem Item
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(body, &newItem)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		newItem.ID = nextID
		nextID++
		items = append(items, newItem)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newItem)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// itemHandler handles requests to /items/{id}.
func itemHandler(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	if len(segments) < 3 {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}
	idStr := segments[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	// Find the item by its ID.
	index := -1
	for i, item := range items {
		if item.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		// Return the specified item.
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items[index])
	case http.MethodDelete:
		// Delete the specified item.
		items = append(items[:index], items[index+1:]...)
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
