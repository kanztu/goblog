package md

import (
	"errors"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

var BootStrapMap = map[string]string{
	"<table>": "<table class=\"table\">",
	"<pre>":   "<pre class=\"prettyprint\">",
	"<code>":  "<code class=\"prettyprint\">",
}

func FetchMDToHtml(fname string) (string, error) {
	md, err := os.ReadFile(fname)
	if err != nil {
		return "", errors.New("Unable to load markdown file")
	}
	extensions := parser.CommonExtensions | parser.Tables | parser.FencedCode | parser.Autolink
	parser := parser.NewWithExtensions(extensions)
	output := string(markdown.ToHTML(md, parser, nil))

	for k, v := range BootStrapMap {
		output = strings.ReplaceAll(output, k, v)
	}

	return output, nil
}
