package main

import (
	"io"
	"os"

	"github.com/nao1215/markdown"
	"github.com/nao1215/markdown/mermaid/flowchart"
)

func main() {
	// Flow Chart (Graph)
	fc := flowchart.NewFlowchart(
		io.Discard,
		flowchart.WithTitle("Flowchart"),
		flowchart.WithOrientalLeftToRight(),
	).
		NodeWithText("A", "Node A").
		String()

	// Output
	markdown.NewMarkdown(os.Stdout).
		CodeBlocks(markdown.SyntaxHighlightMermaid, fc).
		Build()
}
