package main

import (
	"slices"
	"testing"
)

func TestReadDir(t *testing.T) {
	files := readDir()
	mf := "Makefile"

	if !slices.Contains(files, mf) {
		t.Errorf("Makefile not found in dir slice")
	}
}
