package main

import (
	"fmt"
	"os"
	"github.com/unappendixed/recipetomd/schema"
)

func main() {
	data, err := os.ReadFile("full-schema.json")
    if err != nil {
        panic(err)
    }
    
    recipe, err := schema.ParseFromStructuredData(data)
    if err != nil {
        panic(err)
    }

    fmt.Println(recipe.ToMarkdown())

}

