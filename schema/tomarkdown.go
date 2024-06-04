package schema

import (
	"fmt"
	"html"
	"strings"
)

type MDBuilder string

func (md *MDBuilder) WithIndent(level int, s string) {
	*md += MDBuilder(fmt.Sprintf("%s%s", strings.Repeat("  ", level), s))
}

func (r *Recipe) ToMarkdown() string {
	s := fmt.Sprintf("# %s\n\r\n", r.Name)

	if r.Authors[0].Name != "" {
		s += fmt.Sprintf("by %s\n\r\n", r.Authors[0].Name)
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

	s = html.UnescapeString(s)

	return s

}

func (r *Recipe) ToMarkdownLS() string {
	s := MDBuilder(fmt.Sprintf("- # %s\n", r.Name))

	if r.Authors[0].Url == "" {
		s.WithIndent(1, fmt.Sprintf("- by %s\n", r.Authors[0].Name))
	} else {
		s.WithIndent(1, fmt.Sprintf("- by [%s](%s)\n", r.Authors[0].Name, r.Authors[0].Url))
	}
	s.WithIndent(1, fmt.Sprintf("- <%s>\n", r.Url))
	s.WithIndent(1, fmt.Sprintf("- %s\n", r.Description))

	s.WithIndent(1, "- ## Ingredients\n")
	for _, v := range r.Ingredients {
		s.WithIndent(2, fmt.Sprintf("- %s\n", v))
	}

	s.WithIndent(1, "- ## Instructions\n")
	for _, v := range r.Instructions {
		s.WithIndent(2, fmt.Sprintf("- %s\n", v.Text))
		s.WithIndent(2, "logseq.order-list-type:: number\n")
	}

	return html.UnescapeString(string(s))
}
