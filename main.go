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
		fmt.Fprintf(os.Stderr, "Usage: sgviz <json file>\n")
		return
	}
	fileName := os.Args[1]
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read file: %v\n", err)
		return
	}
	var res CLIResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid json: %v\n", err)
		return
	}
	sgs := res.SecurityGroups

	// Generate Markdown
	GenerateMarkDown(os.Stdout, sgs)
}

func GenerateMarkDown(writer io.Writer, sgs []SecurityGroup) {
	// Create Graph
	var g *Graph = NewGraph()

	for _, sg := range sgs {
		g.AddNode(CreateSgNode(&sg))

		for _, ipPerm := range sg.IPPermissions {
			for _, ipRange := range ipPerm.IPRanges {
				g.AddNode(CreateIPNode(&ipRange))
				g.AddEdge(CreatePermEdge(ipRange.CidrIP, sg.GroupID, &ipPerm))
			}

			for _, userIDGroupPair := range ipPerm.UserIDGroupPairs {
				g.AddNode(CreateUserIDGroupPairNode(&userIDGroupPair))
				g.AddEdge(CreatePermEdge(userIDGroupPair.GroupID, sg.GroupID, &ipPerm))
			}

			for _, prefixListId := range ipPerm.PrefixListIds {
				g.AddNode(CreatePrefixNode(&prefixListId))
				g.AddEdge(CreatePermEdge(prefixListId.PrefixListID, sg.GroupID, &ipPerm))
			}

			for _, ipv6Range := range ipPerm.Ipv6Ranges {
				g.AddNode(CreateIpv6Node(&ipv6Range))
				g.AddEdge(CreatePermEdge(ipv6Range.CidrIpv6, sg.GroupID, &ipPerm))
			}
		}
	}

	// Create Flow Chart
	fc := flowchart.NewFlowchart(
		io.Discard,
		flowchart.WithOrientalLeftToRight(),
	)
	for _, n := range g.nodes {
		fc.NodeWithText(n.id, n.text)
	}
	for _, e := range g.edges {
		fc.LinkWithArrowHeadAndText(e.from, e.to, e.text)
	}

	// Generate Output
	markdown.NewMarkdown(os.Stdout).
		CodeBlocks(markdown.SyntaxHighlightMermaid, fc.String()).
		Build()
}
