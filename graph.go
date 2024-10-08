package main

import (
	"fmt"
	"strings"
)

type Node struct {
	id   string
	text string
}
type Edge struct {
	from string
	to   string
	text string
}

type Graph struct {
	nodes map[string]*Node
	edges []*Edge
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]*Node),
		edges: make([]*Edge, 0),
	}
}

func (g *Graph) AddNode(n *Node) {
	if _, ok := g.nodes[n.id]; !ok {
		g.nodes[n.id] = n
	}
}

func (g *Graph) AddEdge(e *Edge) {
	g.edges = append(g.edges, e)
}

func CreateSgNode(sg *SecurityGroup) *Node {
	return &Node{
		id:   sg.GroupID,
		text: fmt.Sprintf("%v:%v\n(%v)", sg.GroupID, sg.GroupName, sg.Description),
	}
}

func CreateIPNode(ip *IPRange) *Node {
	return &Node{
		id:   ip.CidrIP,
		text: fmt.Sprintf("%v\n(%v)", ip.CidrIP, ip.Description),
	}
}

func CreateUserIDGroupPairNode(userIDGroupPair *UserIDGroupPair) *Node {
	return &Node{
		id:   userIDGroupPair.GroupID,
		text: fmt.Sprintf("%v\n(%v)", userIDGroupPair.GroupID, userIDGroupPair.UserID),
	}
}

func CreatePrefixNode(prefix *PrefixListId) *Node {
	return &Node{
		id:   prefix.PrefixListID,
		text: fmt.Sprintf("%v\n(%v)", prefix.PrefixListID, prefix.Description),
	}
}

func CreateIpv6Node(ipv6 *Ipv6Range) *Node {
	return &Node{
		id:   ipv6.CidrIpv6,
		text: fmt.Sprintf("%v\n(%v)", ipv6.CidrIpv6, ipv6.Description),
	}
}

func CreatePermEdge(from string, to string, ipPerm *IPPermission) *Edge {
	var textArray []string
	if ipPerm.IPProtocol == "-1" {
		textArray = append(textArray, "All Protocols")
	} else {
		textArray = append(textArray, strings.ToUpper(ipPerm.IPProtocol))
	}

	if ipPerm.FromPort == 0 && ipPerm.ToPort == 0 {
		textArray = append(textArray, "All Ports")
	} else if ipPerm.FromPort == ipPerm.ToPort {
		textArray = append(textArray, fmt.Sprintf("%v", ipPerm.FromPort))
	} else {
		textArray = append(textArray, fmt.Sprintf("%v-%v", ipPerm.FromPort, ipPerm.ToPort))
	}
	return &Edge{
		from: from,
		to:   to,
		text: strings.Join(textArray, " "),
	}
}
