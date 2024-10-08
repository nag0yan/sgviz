package main

import (
	"fmt"
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

func CreatePermEdge(from string, to string, fromPort int, toPort int) *Edge {
	return &Edge{
		from: from,
		to:   to,
		text: fmt.Sprintf("%v-%v", fromPort, toPort),
	}
}
