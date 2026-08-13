package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func helloWorld() string {
	return "hello world"
}

func downloadImage(url string, fname string, ch chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- "Error querying image url"
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ch <- fmt.Sprintf("Returned status code: %v", resp.StatusCode)
		return
	}

	imgF, err := os.Create(fname)
	if err != nil {
		ch <- "Error opening file"
		return
	}
	defer imgF.Close()

	_, err = io.Copy(imgF, resp.Body)
	if err != nil {
		ch <- "Error copying responde body into file"
		return
	}

	ch <- "Successfully downloaded image"
}

func main() {
	urls := []string{
		"https://go.dev/images/gophers/blue.svg",
		"https://go.dev/images/gophers/motorcycle.svg",
		"https://go.dev/images/gophers/machine.svg",
		"https://go.dev/images/gophers/skateboarding.svg",
	}

	ch := make(chan string)
	for i, url := range urls {
		go downloadImage(url, fmt.Sprintf("./data/downloaded_image_%v.svg", i), ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
}
