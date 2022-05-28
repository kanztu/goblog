package md

import (
	"errors"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func FetchMDToHtml(fname string) (string, error) {
	md, err := os.ReadFile(fname)
	if err != nil {
		return "", errors.New("Unable to load markdown file")
	}
	extensions := parser.CommonExtensions | parser.Tables | parser.FencedCode
	parser := parser.NewWithExtensions(extensions)
	output := markdown.ToHTML(md, parser, nil)
	return string(output), nil
}
