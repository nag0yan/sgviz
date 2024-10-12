package renderer_test

import (
	"bytes"
	"testing"

	"github.com/nag0yan/sgviz/internal/graph"
	"github.com/nag0yan/sgviz/internal/renderer"
)

func TestGenerateMarkDown(t *testing.T) {
	g := graph.NewGraph()
	g.AddNode(&graph.Node{
		Id:   "id",
		Text: "text",
	})
	g.AddEdge(&graph.Edge{
		From: "from",
		To:   "to",
		Text: "text",
	})

	buf := new(bytes.Buffer)

	err := renderer.GenerateMarkDown(buf, g)
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}
	if buf.String() == "" {
		t.Errorf("got empty, want not empty")
	}
}
