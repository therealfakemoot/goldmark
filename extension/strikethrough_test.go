package extension

import (
	"testing"

	"github.com/therealfakemoot/goldmark"
	"github.com/therealfakemoot/goldmark/renderer/html"
	"github.com/therealfakemoot/goldmark/testutil"
)

func TestStrikethrough(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			Strikethrough,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/strikethrough.txt", t, testutil.ParseCliCaseArg()...)
}
