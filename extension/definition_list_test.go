package extension

import (
	"testing"

	"github.com/therealfakemoot/goldmark"
	"github.com/therealfakemoot/goldmark/renderer/html"
	"github.com/therealfakemoot/goldmark/testutil"
)

func TestDefinitionList(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			DefinitionList,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/definition_list.txt", t, testutil.ParseCliCaseArg()...)
}
