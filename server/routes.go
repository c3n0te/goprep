package main

import (
	"encoding/json"
	"net/http"
	"prep/api"
)

var items = map[string]api.Item{
	"1": {ID: "1", Name: "Laptop", Price: 1200},
	"2": {ID: "2", Name: "Mouse", Price: 40},
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func getItemById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	item, exists := items[id]
	if !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem api.Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	items[newItem.ID] = newItem
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}
