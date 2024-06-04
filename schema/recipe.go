package schema

// "@context": "https://schema.org/",
// "@type": "Recipe",

import (
	"encoding/json"
	"fmt"
)

type StructuredData struct {
    Graph []SchemaType `json:"@graph"`
}

type SchemaType struct {
    Type string `json:"@type"`
    Data map[string]any `json:"-"`
}

type Recipe struct {
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	Author       RecipeAuthor        `json:"author"`
	Url          string              `json:"url"`
	Image        []string            `json:"image"`
	Ingredients  []RecipeIngredient  `json:"recipeIngredient"`
	Instructions []RecipeInstruction `json:"recipeInstructions"`
}

type RecipeIngredient string

type RecipeInstruction struct {
	Text string `json:"text"`
}

type RecipeAuthor struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func ParseFromStructuredData(in []byte) (*Recipe, error) {
    sd := StructuredData{}

    err := json.Unmarshal(in, &sd)
    if err != nil {
        panic(err)
    }

    graphMap := struct{Graph []json.RawMessage `json:"@graph"`}{}

    err = json.Unmarshal(in, &graphMap)
    if err != nil {
        panic(err)
    }

    var rawRecipe json.RawMessage
    for i, v := range sd.Graph {
        if v.Type == "Recipe" {
            rawRecipe = graphMap.Graph[i]
        }
    }

    return New(rawRecipe)
}

func New(in []byte) (*Recipe, error) {
    recipe := &Recipe{}
	err := json.Unmarshal(in, recipe)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal string: %w", err)
	}

	return recipe, nil
}
