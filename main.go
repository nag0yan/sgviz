package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/nag0yan/sgviz/internal/model"
	"github.com/nag0yan/sgviz/internal/graph"
	"github.com/nao1215/markdown"
	"github.com/nao1215/markdown/mermaid/flowchart"
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

	// Generate Markdown
	GenerateMarkDown(os.Stdout, sgs)
}

func GenerateMarkDown(writer io.Writer, sgs []model.SecurityGroup) {
	// Create Graph
	var g *graph.Graph = graph.NewGraph()

	for _, sg := range sgs {
		g.AddNode(graph.CreateSgNode(&sg))

		for _, ipPerm := range sg.IPPermissions {
			for _, ipRange := range ipPerm.IPRanges {
				g.AddNode(graph.CreateIPNode(&ipRange))
				g.AddEdge(graph.CreatePermEdge(ipRange.CidrIP, sg.GroupID, &ipPerm))
			}

			for _, userIDGroupPair := range ipPerm.UserIDGroupPairs {
				g.AddNode(graph.CreateUserIDGroupPairNode(&userIDGroupPair))
				g.AddEdge(graph.CreatePermEdge(userIDGroupPair.GroupID, sg.GroupID, &ipPerm))
			}

			for _, prefixListId := range ipPerm.PrefixListIds {
				g.AddNode(graph.CreatePrefixNode(&prefixListId))
				g.AddEdge(graph.CreatePermEdge(prefixListId.PrefixListID, sg.GroupID, &ipPerm))
			}

			for _, ipv6Range := range ipPerm.Ipv6Ranges {
				g.AddNode(graph.CreateIpv6Node(&ipv6Range))
				g.AddEdge(graph.CreatePermEdge(ipv6Range.CidrIpv6, sg.GroupID, &ipPerm))
			}
		}
	}

	// Create Flow Chart
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
	markdown.NewMarkdown(os.Stdout).
		CodeBlocks(markdown.SyntaxHighlightMermaid, fc.String()).
		Build()
}
