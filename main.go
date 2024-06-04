package main

import (
	"flag"
	"fmt"
	"github.com/unappendixed/recipetomd/schema"
)

func main() {

    flag.Parse()

    url := flag.Arg(0)

    jsonld, err := ScrapeUrl(url)
    if err != nil {
        panic(err)
    }

    recipe, err := schema.ParseFromStructuredData([]byte(jsonld))
    if err != nil {
        panic(err)
    }

    if recipe.Url == "" {
        recipe.Url = url
    }

    if *MarkdownFlavor {
        fmt.Println(recipe.ToMarkdownLS())
    } else {
        fmt.Println(recipe.ToMarkdown())
    }

}

