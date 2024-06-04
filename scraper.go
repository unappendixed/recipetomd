package main

import (
	"strings"

	"github.com/gocolly/colly"
)

func ScrapeUrl(url string) (string, error) {
	c := colly.NewCollector()

	var schema string
	c.OnHTML("script[type=\"application/ld+json\"]", func(h *colly.HTMLElement) {
		if strings.Contains(h.Text, "schema.org") {
			schema = h.Text
		}
	})

	err := c.Visit(url)
	if err != nil {
		return "", err
	}

	return schema, nil

}
