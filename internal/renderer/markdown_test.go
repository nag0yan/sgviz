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
	g.AddNode("sg-12345678", "text")
	g.AddEdge("xx.xx.xx.xx", "to", "text")

	buf := new(bytes.Buffer)

	err := renderer.GenerateMarkDown(buf, g)
	output := buf.String()
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}
	if output == "" {
		t.Errorf("got empty, want not empty")
	}
	c := []string{"```", "mermaid", "sg-12345678", "xx.xx.xx.xx"}
	for _, v := range c {
		if !strings.Contains(output, v) {
			t.Errorf("got %v, want to contain %v", output, v)
		}
	}
}
