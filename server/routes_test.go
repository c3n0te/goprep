package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"prep/api"
	"testing"
)

func TestRoutes(t *testing.T) {
	mux := newRouter()
	server := httptest.NewServer(mux)
	defer server.Close()

	t.Run("GET /items returns all items", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/items")
		if err != nil {
			t.Fatalf("Failed to send GET request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected 200 status code, got %v", resp.StatusCode)
		}

		var returnedItems map[string]api.Item
		if err := json.NewDecoder(resp.Body).Decode(&returnedItems); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if len(returnedItems) != 2 {
			t.Errorf("Expected 2 items, got %d", len(returnedItems))
		}
	})

	t.Run("GET /items/1 returns specific item", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/items/1")
		if err != nil {
			t.Fatalf("Failed to send GET request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected 200 status code, got %v", resp.StatusCode)
		}

		var item api.Item
		if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if item.Name != "Laptop" {
			t.Errorf("Expected to receive 'Laptop' received %v", item.Name)
		}
	})

	t.Run("POST /items successfully creates new record", func(t *testing.T) {
		newItem := api.Item{
			ID:    "3",
			Name:  "Keyboard",
			Price: 80,
		}

		jsonData, _ := json.Marshal(newItem)
		resp, err := http.Post(server.URL+"/items", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatalf("Failed to send POST request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("Expected 200 status code, got %v", resp.StatusCode)
		}

		var item api.Item
		json.NewDecoder(resp.Body).Decode(&item)

		if item.ID != "3" || item.Name != "Keyboard" {
			t.Errorf("Returned incorrect item, got: %v", item)
		}
	})

	t.Run("GET /items/99 returns 404 error code", func(*testing.T) {
		resp, err := http.Get(server.URL + "/items/99")
		if err != nil {
			t.Fatalf("Failed to send GET request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Expected 404, got: %v", resp.StatusCode)
		}
	})
}
