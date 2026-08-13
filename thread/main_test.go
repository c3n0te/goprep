package main

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	s := helloWorld()
	expected := "hello world"
	if s != expected {
		t.Errorf("Expected string did not match")
	}
}

func TestDownloadImage(t *testing.T) {
	ch := make(chan string)
	go downloadImage("https://go.dev/images/gophers/skateboarding.svg", fmt.Sprintf("../data/downloaded_image_%v.svg", 1), ch)
	successStr := <-ch
	if successStr != "Successfully downloaded image" {
		t.Errorf("Failed to download image")
	}
}
