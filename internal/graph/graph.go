package graph

import (
	"fmt"
	"strings"
	"github.com/nag0yan/sgviz/internal/model"
)

type Node struct {
	Id   string
	Text string
}
type Edge struct {
	From string
	To   string
	Text string
}

type Graph struct {
	nodes map[string]*Node
	edges []*Edge
}

func (g *Graph) GetNodes() map[string]*Node {
	return g.nodes
}

func (g *Graph) GetEdges() []*Edge {
	return g.edges
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]*Node),
		edges: make([]*Edge, 0),
	}
}

func (g *Graph) AddNode(n *Node) {
	if _, ok := g.nodes[n.Id]; !ok {
		g.nodes[n.Id] = n
	}
}

func (g *Graph) AddEdge(e *Edge) {
	g.edges = append(g.edges, e)
}

func CreateSgNode(sg *model.SecurityGroup) *Node {
	return &Node{
		Id:   sg.GroupID,
		Text: fmt.Sprintf("%v:%v\n(%v)", sg.GroupID, sg.GroupName, sg.Description),
	}
}

func CreateIPNode(ip *model.IPRange) *Node {
	return &Node{
		Id:   ip.CidrIP,
		Text: fmt.Sprintf("%v\n(%v)", ip.CidrIP, ip.Description),
	}
}

func CreateUserIDGroupPairNode(userIDGroupPair *model.UserIDGroupPair) *Node {
	return &Node{
		Id:   userIDGroupPair.GroupID,
		Text: fmt.Sprintf("%v\n(%v)", userIDGroupPair.GroupID, userIDGroupPair.UserID),
	}
}

func CreatePrefixNode(prefix *model.PrefixListId) *Node {
	return &Node{
		Id:   prefix.PrefixListID,
		Text: fmt.Sprintf("%v\n(%v)", prefix.PrefixListID, prefix.Description),
	}
}

func CreateIpv6Node(ipv6 *model.Ipv6Range) *Node {
	return &Node{
		Id:   ipv6.CidrIpv6,
		Text: fmt.Sprintf("%v\n(%v)", ipv6.CidrIpv6, ipv6.Description),
	}
}

func CreatePermEdge(from string, to string, ipPerm *model.IPPermission) *Edge {
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
		From: from,
		To:   to,
		Text: strings.Join(textArray, " "),
	}
}
