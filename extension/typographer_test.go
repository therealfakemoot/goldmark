package extension

import (
	"testing"

	"github.com/therealfakemoot/goldmark"
	"github.com/therealfakemoot/goldmark/renderer/html"
	"github.com/therealfakemoot/goldmark/testutil"
)

func TestTypographer(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			Typographer,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/typographer.txt", t, testutil.ParseCliCaseArg()...)
}
