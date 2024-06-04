package main

import (
	"strings"

	"github.com/gocolly/colly"
)

func ScrapeUrl(url string) ([]string, error) {
	c := colly.NewCollector()

    schemas := []string{}
	c.OnHTML("script[type=\"application/ld+json\"]", func(h *colly.HTMLElement) {
		if strings.Contains(h.Text, "schema.org") {
			schemas = append(schemas, h.Text)
		}
	})

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return schemas, nil

}
