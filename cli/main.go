package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
)

func parseArgs() string {
	first := flag.String("dir", "directory flag", "enter directory path you'd like to list")
	flag.Parse()
	return *first
}

func readDir() []string {
	dir := parseArgs()
	entries, err := os.ReadDir(dir)
	if err != nil {
		slog.Error("Failed to read directory", "error", err)
	}

	fnames := []string{}
	for _, entry := range entries {
		fname := entry.Name()
		fnames = append(fnames, fname)
	}

	return fnames
}

func main() {
	files := readDir()
	for _, fname := range files {
		fmt.Println(fname)
	}
}
