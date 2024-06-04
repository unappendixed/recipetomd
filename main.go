package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/unappendixed/recipetomd/schema"
)

func main() {

    flag.Parse()

    url := flag.Arg(0)

    jsonld, err := ScrapeUrl(url)
    if err != nil {
        fmt.Println("Couldn't scrape webpage, exiting...")
        os.Exit(1)
    }

    recipes := []schema.Recipe{}
    for _, v := range jsonld {
        recipe, err := schema.ParseFromStructuredData([]byte(v))
        if err == nil {
            recipes = append(recipes, *recipe)
        }
    }

    if len(recipes) == 0 {
        fmt.Println("Couldn't find any structured recipe data!")
    }

    var recipe schema.Recipe
    if len(recipes) >= 0 {
        recipe = recipes[0]
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

