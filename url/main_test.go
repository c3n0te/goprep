package main

import "testing"

func TestHelloWorld(t *testing.T) {
	s := helloWorld()
	expected := "hello world"
	if s != expected {
		t.Errorf("Strings do not match")
	}
}
