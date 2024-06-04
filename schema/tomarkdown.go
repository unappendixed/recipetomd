package schema

import "fmt"

func (r *Recipe) ToMarkdown() string {
    s := fmt.Sprintf("# %s\n\r\n", r.Name)

    if r.Author.Name != "" {
        s += fmt.Sprintf("by %s\n\r\n", r.Author.Name)
    }

    s += fmt.Sprintf("<%s>\n\n", r.Url)

    if r.Description != "" {
        s += r.Description + "\n\n"
    }

    s += "### Ingredients\n\n"
    for _, v := range r.Ingredients {
        s += fmt.Sprintf("- %s\n", v)
    }
    s += "\n"

    s += "### Instructions\n\n"
    for i, v := range r.Instructions {
        s += fmt.Sprintf("%d. %s\n", i+1, v.Text)
    }
    s += "\n"

    return s

}
