package main

import (
	"log/slog"
	"net/http"
	"sync"
)

type URL struct {
	Url string `json:"url"`
}

type DB struct {
	mu        sync.RWMutex
	urls      map[int]URL
	shortUrls map[URL]URL
	stats     map[URL]int
	nextId    int
}

var db = DB{
	urls:      make(map[int]URL),
	shortUrls: make(map[URL]URL),
	stats:     make(map[URL]int),
	nextId:    123456,
}

func helloWorld() string {
	return "hello world"
}

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /shorten", shortenUrl)
	return mux
}

func main() {
	mux := newRouter()

	slog.Info("Server running on port :8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		slog.Error("Failed to listen on port 8000")
		panic(err)
	}
}
