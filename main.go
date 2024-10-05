package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/nao1215/markdown"
	"github.com/nao1215/markdown/mermaid/flowchart"
)

func main() {
	// Load json
	if len(os.Args) < 2 {
		fmt.Println("Usage: sgviz <json file>")
		return
	}
	fileName := os.Args[1]
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return
	}
	var res CLIResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		fmt.Printf("Invalid json: %v\n", err)
		return
	}
	sgs := res.SecurityGroups

	// Create Flow Chart (Graph)
	fc := flowchart.NewFlowchart(
		io.Discard,
		flowchart.WithOrientalLeftToRight(),
	)

	for _, sg := range sgs {
		fc.NodeWithText(sg.GroupID, fmt.Sprintf("%v\n(%v)", sg.GroupID, sg.GroupName))
	}

	// Generate Output
	markdown.NewMarkdown(os.Stdout).
		CodeBlocks(markdown.SyntaxHighlightMermaid, fc.String()).
		Build()
}
