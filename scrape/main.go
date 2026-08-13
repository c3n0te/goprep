package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("snow-forecast.com", "www.snow-forecast.com"),
	)

	c.OnResponse(func(r *colly.Response) {
		htmlRaw := string(r.Body)
		fmt.Println(htmlRaw)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error visiting %s: %v", r.Request.URL, err)
	})

	if err := c.Visit("https://www.snow-forecast.com/resorts/HemlockResort/6day/mid"); err != nil {
		log.Fatal(err)
	}
}
