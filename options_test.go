package goldmark_test

import (
	"testing"

	. "github.com/therealfakemoot/goldmark"
	"github.com/therealfakemoot/goldmark/parser"
	"github.com/therealfakemoot/goldmark/testutil"
)

func TestAttributeAndAutoHeadingID(t *testing.T) {
	markdown := New(
		WithParserOptions(
			parser.WithAttribute(),
			parser.WithAutoHeadingID(),
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/options.txt", t, testutil.ParseCliCaseArg()...)
}
