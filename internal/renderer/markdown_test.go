package renderer_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/nag0yan/sgviz/internal/graph"
	"github.com/nag0yan/sgviz/internal/renderer"
)

func TestGenerateMarkDown(t *testing.T) {
	g := graph.NewGraph()
	g.AddNode(&graph.Node{
		Id:   "sg-12345678",
		Text: "text",
	})
	g.AddEdge(&graph.Edge{
		From: "xx.xx.xx.xx",
		To:   "to",
		Text: "text",
	})

	buf := new(bytes.Buffer)

	err := renderer.GenerateMarkDown(buf, g)
	output := buf.String()
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}
	if output == "" {
		t.Errorf("got empty, want not empty")
	}
	if !strings.Contains(output, "mermaid") {
		t.Errorf("got %v, want to contain %v", output, "mermaid")
	}
	if !strings.Contains(output, "sg-12345678") {
		t.Errorf("got %v, want to contain %v", output, "sg-12345678")
	}
	if !strings.Contains(output, "xx.xx.xx.xx") {
		t.Errorf("got %v, want to contain %v", output, "xx.xx.xx.xx")
	}
}
