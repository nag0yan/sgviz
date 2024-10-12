package renderer

import (
	"io"

	"github.com/nag0yan/sgviz/internal/graph"
	"github.com/nao1215/markdown"
	"github.com/nao1215/markdown/mermaid/flowchart"
)

func GenerateMarkDown(writer io.Writer, g *graph.Graph) error {
	fc := flowchart.NewFlowchart(
		io.Discard,
		flowchart.WithOrientalLeftToRight(),
	)
	for _, n := range g.GetNodes() {
		fc.NodeWithText(n.Id, n.Text)
	}
	for _, e := range g.GetEdges() {
		fc.LinkWithArrowHeadAndText(e.From, e.To, e.Text)
	}

	// Generate Output
	err := markdown.NewMarkdown(writer).
		CodeBlocks(markdown.SyntaxHighlightMermaid, fc.String()).
		Build()

	return err
}
