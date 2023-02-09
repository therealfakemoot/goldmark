package main

import (
	"bytes"
	"fmt"

	"github.com/therealfakemoot/goldmark"
	"github.com/therealfakemoot/goldmark/extension"
)

func main() {
	source := "[[Regular Link]]"
	md := goldmark.New(
		goldmark.WithExtensions(extension.NewWikilinksExtension(nil)),
	)
	var buf bytes.Buffer
	if err := md.Convert([]byte(source), &buf); err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}
