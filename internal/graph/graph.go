package graph

import (
	"fmt"
	"sort"
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
	nodes       []*Node
	node_exists map[string]bool
	edges       []*Edge
}

func NewGraph() *Graph {
	return &Graph{
		nodes:       make([]*Node, 0),
		edges:       make([]*Edge, 0),
		node_exists: make(map[string]bool),
	}
}

func (g *Graph) GetNodes() []*Node {
	return g.nodes
}

func (g *Graph) GetEdges() []*Edge {
	return g.edges
}

func (g *Graph) IfNodeExist(id string) bool {
	return g.node_exists[id]
}

func (g *Graph) AddNode(id string, text string) {
	if !g.node_exists[id] {
		g.nodes = append(g.nodes, &Node{
			Id:   id,
			Text: text,
		})
		g.node_exists[id] = true
	}
}

func (g *Graph) AddEdge(from string, to string, text string) {
	g.edges = append(g.edges, &Edge{
		From: from,
		To:   to,
		Text: text,
	})
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

func GenerateGraph(sgs []model.SecurityGroup) (*Graph, error) {
	var g *Graph = NewGraph()

	for _, sg := range sgs {
		sgn := CreateSgNode(&sg)
		g.AddNode(sgn.Id, sgn.Text)

		for _, ipPerm := range sg.IPPermissions {
			for _, ipRange := range ipPerm.IPRanges {
				n := CreateIPNode(&ipRange)
				g.AddNode(n.Id, n.Text)
				e := CreatePermEdge(ipRange.CidrIP, sg.GroupID, &ipPerm)
				g.AddEdge(e.From, e.To, e.Text)
			}

			for _, userIDGroupPair := range ipPerm.UserIDGroupPairs {
				n := CreateUserIDGroupPairNode(&userIDGroupPair)
				g.AddNode(n.Id, n.Text)
				e := CreatePermEdge(userIDGroupPair.GroupID, sg.GroupID, &ipPerm)
				g.AddEdge(e.From, e.To, e.Text)
			}

			for _, prefixListId := range ipPerm.PrefixListIds {
				n := CreatePrefixNode(&prefixListId)
				g.AddNode(n.Id, n.Text)
				e := CreatePermEdge(prefixListId.PrefixListID, sg.GroupID, &ipPerm)
				g.AddEdge(e.From, e.To, e.Text)
			}

			for _, ipv6Range := range ipPerm.Ipv6Ranges {
				n := CreateIpv6Node(&ipv6Range)
				g.AddNode(n.Id, n.Text)
				e := CreatePermEdge(ipv6Range.CidrIpv6, sg.GroupID, &ipPerm)
				g.AddEdge(e.From, e.To, e.Text)
			}
		}
	}

	return g, nil
}

func (g *Graph) AggregateNodes() *Graph {
	// 1. Create edge map for aggregation
	// Each node has edge set such as {target: n, isout: true, text: "TCP 80"}
	// 2. Compare each node and aggregate if they have same edge set
	// 3. Return new graph with aggregated nodes

	// 1. Create edge map for aggregation
	edgeMap := make(map[string][]*DirectedEdge)
	for _, edge := range g.GetEdges() {
		edgeMap[edge.From] = append(edgeMap[edge.From], &DirectedEdge{
			Target: edge.To,
			IsOut:  true,
			Text:   edge.Text,
		})
		edgeMap[edge.To] = append(edgeMap[edge.To], &DirectedEdge{
			Target: edge.From,
			IsOut:  false,
			Text:   edge.Text,
		})
	}

	// 2. Compare each node and aggregate if they have same edge set
	newg := NewGraph()
	agged := make(map[string]bool)
	for i, node := range g.GetNodes() {
		if agged[node.Id] {
			continue
		}
		id := node.Id
		text := node.Text
		edges := edgeMap[node.Id]

		for j, target := range g.GetNodes() {
			// Compare edge set
			e1 := edgeMap[node.Id]
			e2 := edgeMap[target.Id]

			if i > j {
				continue
			}
			if agged[target.Id] {
				continue
			}
			if node.Id == target.Id {
				continue
			}
			if len(edgeMap[node.Id]) != len(edgeMap[target.Id]) {
				continue
			}

			if EqualEdgeSet(e1, e2) {
				text += fmt.Sprintf("\n%v", target.Text)
				agged[target.Id] = true
			}
		}
		for _, edge := range edges {
			if agged[edge.Target] {}
		}
		newg.AddNode(id, text)
	}
	for _, edge := range g.GetEdges() {
		if agged[edge.From] || agged[edge.To] {
			continue
		}
		newg.AddEdge(edge.From, edge.To, edge.Text)
	}

	return newg
}

type DirectedEdge struct {
	Target string
	IsOut  bool
	Text   string
}

type DirectedEdges []*DirectedEdge

func (e DirectedEdges) Len() int {
	return len(e)
}
func (e DirectedEdges) Less(i, j int) bool {
	return (e[i].Target < e[j].Target) || (!e[i].IsOut) || (e[i].Text < e[j].Text)
}
func (e DirectedEdges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
func SortDirectedEdges(edges []*DirectedEdge) {
	sort.Sort(DirectedEdges(edges))
}

func EqualDirectedEdge(e1 *DirectedEdge, e2 *DirectedEdge) bool {
	return e1.Target == e2.Target && e1.IsOut == e2.IsOut && e1.Text == e2.Text
}

func EqualEdgeSet(e1 DirectedEdges, e2 DirectedEdges) bool {
	if len(e1) != len(e2) {
		return false
	}
	SortDirectedEdges(e1)
	SortDirectedEdges(e2)
	for i, e := range e1 {
		if !EqualDirectedEdge(e, e2[i]) {
			return false
		}
	}
	return true
}

func (e *DirectedEdge) String() string {
	return fmt.Sprintf("(%v, %v, %v)", e.Target, e.IsOut, e.Text)
}

func (g *Graph) GetNode(id string) *Node {
	for _, node := range g.GetNodes() {
		if node.Id == id {
			return node
		}
	}
	return nil
}

func (g *Graph) GetNodeCount() int {
	return len(g.GetNodes())
}

func (g *Graph) GetEdgeCount() int {
	return len(g.GetEdges())
}
