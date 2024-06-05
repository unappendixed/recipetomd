package schema

// "@context": "https://schema.org/",
// "@type": "Recipe",

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

type StructuredData struct {
	Graph []SchemaType `json:"@graph"`
}

type SchemaType struct {
	Type SchemaTypeArray `json:"@type"`
}

type SchemaTypeArray []string

func (sta *SchemaTypeArray) UnmarshalJSON(data []byte) error {
	list := []string{}

	err := json.Unmarshal(data, &list)
	if err == nil {
		*sta = SchemaTypeArray(list)
		return nil
	}

	var str string
	err = json.Unmarshal(data, &str)
	if err == nil {
		*sta = SchemaTypeArray([]string{str})
		return nil
	}

	return err
}

type Recipe struct {
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	Authors      RecipeAuthorList   `json:"author"`
	Url          string             `json:"url"`
	Ingredients  []RecipeIngredient `json:"recipeIngredient"`
	Instructions RecipeHowToSection `json:"recipeInstructions"`
}

type RecipeHowToSection []RecipeInstruction

func (htd *RecipeHowToSection) UnmarshalJSON(data []byte) error {

	type wrappedInstructions struct {
		Items []RecipeInstruction `json:"itemListElement"`
	}

	type manyWrappedInstructions []wrappedInstructions

	unwrapped := []RecipeInstruction{}
    err := json.Unmarshal(data, &unwrapped)
	if err == nil && len(unwrapped) > 0 {
		*htd = unwrapped
		return nil
	}

	wrapped := wrappedInstructions{}

	err = json.Unmarshal(data, &wrapped)
	if err == nil && len(wrapped.Items) > 0 {
		*htd = wrapped.Items
		return nil
	}

	manyWrapped := manyWrappedInstructions{}

	err = json.Unmarshal(data, &manyWrapped)
	if err == nil && len(manyWrapped) > 0 {
		items := []RecipeInstruction{}
		for _, v := range manyWrapped {
			items = slices.Concat(items, v.Items)
		}
        if len(items) == 0 {
            fmt.Println("Failed to parse recipe instructions")
            os.Exit(1)
        }
		*htd = items
		return nil
	}



	fmt.Printf("Failed to parse recipe: %s\n", err.Error())
	os.Exit(1)
	return nil
}

type RecipeIngredient string

type RecipeInstruction struct {
	Text string `json:"text"`
}

type RecipeAuthorList []RecipeAuthor

func (ral *RecipeAuthorList) UnmarshalJSON(data []byte) error {
	list := []RecipeAuthor{}

	err := json.Unmarshal(data, &list)
	if err == nil {
		*ral = list
		return nil
	}

	single := RecipeAuthor{}

	err = json.Unmarshal(data, &single)
	if err == nil {
		*ral = []RecipeAuthor{single}
		return nil
	}

	return err
}

type RecipeAuthor struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func ParseFromStructuredData(in []byte) (*Recipe, error) {
	sd := StructuredData{}

	err := json.Unmarshal(in, &sd)
	if err != nil || len(sd.Graph) == 0 {
		return ParseFromSDList(in)
	}

	graphMap := struct {
		Graph []json.RawMessage `json:"@graph"`
	}{}

	err = json.Unmarshal(in, &graphMap)
	if err != nil {
		fmt.Println("Failed to parse structured data, exiting...")
		os.Exit(1)
	}

	var rawRecipe json.RawMessage
	for i, v := range sd.Graph {
		if slices.Contains(v.Type, "Recipe") {
			rawRecipe = graphMap.Graph[i]
		}
	}

	return New(rawRecipe)
}

func ParseFromSDList(in []byte) (*Recipe, error) {
	sd := []SchemaType{}

	err := json.Unmarshal(in, &sd)
	if err != nil {
		return New(in)
	}

	entries := []json.RawMessage{}

	err = json.Unmarshal(in, &entries)
	if err != nil {
		fmt.Println("Failed to parse structured data, exiting...")
		os.Exit(1)
	}

	for i, v := range sd {
		if slices.Contains(v.Type, "Recipe") {
			return New([]byte(entries[i]))
		}
	}

	return New(in)
}

func New(in []byte) (*Recipe, error) {
	recipe := &Recipe{}
	err := json.Unmarshal(in, recipe)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal string: %w", err)
	}

	return recipe, nil
}
