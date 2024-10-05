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

	// Create Flow Chart (Graph)
	fc := flowchart.NewFlowchart(
		io.Discard,
		flowchart.WithOrientalLeftToRight(),
	)

	for _, sg := range sgs {
		fc.NodeWithText(sg.GroupID, fmt.Sprintf("%v\n(%v)", sg.GroupID, sg.GroupName))

		for _, ipPerm := range sg.IPPermissions {
			for _, ipRange := range ipPerm.IPRanges {
				fc.NodeWithText(ipRange.CidrIP, fmt.Sprintf("%v\n(%v)", ipRange.CidrIP, ipRange.Description))
				fc.LinkWithArrowHeadAndText(fmt.Sprintf("%v", ipRange.CidrIP), sg.GroupID, fmt.Sprintf("%v-%v", ipPerm.FromPort, ipPerm.ToPort))
			}

			for _, userIDGroupPair := range ipPerm.UserIDGroupPairs {
				fc.LinkWithArrowHeadAndText(fmt.Sprintf("%v", userIDGroupPair.GroupID), sg.GroupID, fmt.Sprintf("%v-%v", ipPerm.FromPort, ipPerm.ToPort))
			}

			for _, prefixListId := range ipPerm.PrefixListIds {
				fc.NodeWithText(prefixListId.PrefixListID, fmt.Sprintf("%v\n(%v)", prefixListId.PrefixListID, prefixListId.Description))
				fc.LinkWithArrowHeadAndText(fmt.Sprintf("%v", prefixListId.PrefixListID), sg.GroupID, fmt.Sprintf("%v-%v", ipPerm.FromPort, ipPerm.ToPort))
			}

			for _, ipv6Range := range ipPerm.Ipv6Ranges {
				fc.NodeWithText(ipv6Range.CidrIpv6, fmt.Sprintf("%v\n(%v)", ipv6Range.CidrIpv6, ipv6Range.Description))
				fc.LinkWithArrowHeadAndText(fmt.Sprintf("%v", ipv6Range.CidrIpv6), sg.GroupID, fmt.Sprintf("%v-%v", ipPerm.FromPort, ipPerm.ToPort))
			}
		}
	}

	// Generate Output
	markdown.NewMarkdown(os.Stdout).
		CodeBlocks(markdown.SyntaxHighlightMermaid, fc.String()).
		Build()
}
