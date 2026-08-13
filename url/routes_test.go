package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	mux := newRouter()
	server := httptest.NewServer(mux)
	defer server.Close()

	t.Run("POST /shorten/https://www.example.com/some/long/url successfully returns shortened url", func(t *testing.T) {
		url := URL{
			Url: "https://www.example.com/some/long/url",
		}

		jsonData, _ := json.Marshal(url)
		resp, err := http.Post(server.URL+"/shorten", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatalf("Failed to post to /shorten: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("Failed to create new shortened url: %v", err)
		}

		var shortenedUrl URL
		json.NewDecoder(resp.Body).Decode(&shortenedUrl)
		if shortenedUrl.Url != "w7f" {
			t.Errorf("Error wrong shortened url")
		}
	})
}
