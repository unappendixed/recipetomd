# Recipe to Markdown

## Overview
Simple Go CLI tool to scrape recipe webpages and convert it into clean, portable markdown.

## Demo
https://github.com/unappendixed/recipetomd/assets/38593737/2ce0a778-4fa3-469e-9ae0-e7850c39c246

## Installation
### Binary
1. Download the appropriate binary for your system from [Releases](github.com/unappendixed/recipetomd/releases/latest)
2. Move the binary file to a folder on the PATH.

### Go
1. Install the [Go toolchain](https://go.dev)
2. Run `go install https://github.com/unappendixed/recipetomd@latest`

### Usage
- Just execute `recipetomd <recipe-url>` to spit out markdown to stdout.
- Redirect to a file:\
  `recipetomd <recipe-url> > recipe.md`
- Redirect to another program (e.g. [Glow](https://github.com/charmbracelet/glow)) \
  `recipetomd <recipe-url> | glow`
- Generate Logseq-flavoured markdown with the `-ls` flag.

