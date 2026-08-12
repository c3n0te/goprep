package main

import (
	"slices"
	"testing"
)

func TestReadDir(t *testing.T) {
	dir := "."
	files := readDir(dir)
	mf := "Makefile"

	if !slices.Contains(files, mf) {
		t.Errorf("Makefile not found in dir slice")
	}
}
