package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nag0yan/sgviz/internal/graph"
	"github.com/nag0yan/sgviz/internal/model"
	"github.com/nag0yan/sgviz/internal/renderer"
)

func main() {
	// Load json
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: sgviz <json file>\n")
		return
	}
	fileName := os.Args[1]
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read file: %v\n", err)
		return
	}
	var res model.CLIResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid json: %v\n", err)
		return
	}
	sgs := res.SecurityGroups

	g, err := graph.GenerateGraph(sgs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate graph: %v\n", err)
		return
	}

	err = renderer.GenerateMarkDown(os.Stdout, g)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate markdown: %v\n", err)
		return
	}
}
