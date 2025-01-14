package extension

import (
	"testing"

	"github.com/therealfakemoot/goldmark"
	"github.com/therealfakemoot/goldmark/renderer/html"
	"github.com/therealfakemoot/goldmark/testutil"
)

func TestTaskList(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			TaskList,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/tasklist.txt", t, testutil.ParseCliCaseArg()...)
}
