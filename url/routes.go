package main

import (
	"encoding/json"
	"net/http"
)

func shortenUrl(w http.ResponseWriter, r *http.Request) {
	var url URL
	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	db.mu.Lock()
	db.nextId++
	db.urls[db.nextId] = url
	db.stats[url] += 1
	shortenedUrlStr := Encode(uint64(db.nextId))
	shortenedUrl := URL{Url: shortenedUrlStr}
	db.shortUrls[shortenedUrl] = url
	db.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shortenedUrl)
}
