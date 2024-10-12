package graph_test

import (
	"testing"

	"github.com/nag0yan/sgviz/internal/graph"
	"github.com/nag0yan/sgviz/internal/model"
)

func TestCreateSgNode(t *testing.T) {
	sg := &model.SecurityGroup{
		GroupID:     "id",
		GroupName:   "name",
		Description: "description",
	}
	got := graph.CreateSgNode(sg)
	want := &graph.Node{
		Id:   "id",
		Text: "id:name\n(description)",
	}
	if *got != *want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCreateIPNode(t *testing.T) {
	ip := &model.IPRange{
		CidrIP:      "cidr",
		Description: "description",
	}
	got := graph.CreateIPNode(ip)
	want := &graph.Node{
		Id:   "cidr",
		Text: "cidr\n(description)",
	}
	if *got != *want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCreateUserIDGroupPairNode(t *testing.T) {
	userIDGroupPair := &model.UserIDGroupPair{
		GroupID: "id",
		UserID:  "user",
	}
	got := graph.CreateUserIDGroupPairNode(userIDGroupPair)
	want := &graph.Node{
		Id:   "id",
		Text: "id\n(user)",
	}
	if *got != *want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCreatePrefixNode(t *testing.T) {
	prefix := &model.PrefixListId{
		PrefixListID: "id",
		Description:  "description",
	}
	got := graph.CreatePrefixNode(prefix)
	want := &graph.Node{
		Id:   "id",
		Text: "id\n(description)",
	}
	if *got != *want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCreateIpv6Node(t *testing.T) {
	ipv6 := &model.Ipv6Range{
		CidrIpv6:    "cidr",
		Description: "description",
	}
	got := graph.CreateIpv6Node(ipv6)
	want := &graph.Node{
		Id:   "cidr",
		Text: "cidr\n(description)",
	}
	if *got != *want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCreatePermEdge(t *testing.T) {
	got := graph.CreatePermEdge("from", "to", &model.IPPermission{
		FromPort:   1,
		ToPort:     2,
		IPProtocol: "tcp",
	})
	want := &graph.Edge{
		From: "from",
		To:   "to",
		Text: "TCP 1-2",
	}
	if *got != *want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCreateSinglePortPermEdge(t *testing.T) {
	got := graph.CreatePermEdge("from", "to", &model.IPPermission{
		FromPort:   1,
		ToPort:     1,
		IPProtocol: "tcp",
	})
	want := &graph.Edge{
		From: "from",
		To:   "to",
		Text: "TCP 1",
	}
	if *got != *want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCreateZeroPortPermEdge(t *testing.T) {
	got := graph.CreatePermEdge("from", "to", &model.IPPermission{
		FromPort:   0,
		ToPort:     0,
		IPProtocol: "tcp",
	})
	want := &graph.Edge{
		From: "from",
		To:   "to",
		Text: "TCP All Ports",
	}
	if *got != *want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestAddNode(t *testing.T) {
	g := graph.NewGraph()
	n := &graph.Node{Id: "id"}
	g.AddNode(n)
	got := g.GetNodes()["id"]
	if got != n {
		t.Errorf("got %v, want %v", got, n)
	}
}

func TestAddEdge(t *testing.T) {
	g := graph.NewGraph()
	e := &graph.Edge{}
	g.AddEdge(e)
	got := g.GetEdges()[0]
	if got != e {
		t.Errorf("got %v, want %v", got, e)
	}
}

func TestGenerateGraph(t *testing.T) {
	sgs := []model.SecurityGroup{
		{
			GroupID:     "id",
			GroupName:   "name",
			Description: "description",
			IPPermissions: []model.IPPermission{
				{
					FromPort:   1,
					ToPort:     2,
					IPProtocol: "tcp",
					IPRanges: []model.IPRange{
						{
							CidrIP:      "10.0.0.0",
							Description: "description",
						},
					},
				},
			},
		},
	}
	type result struct {
		err  error
		ncnt int
		ecnt int
	}
	want := result{
		err:  nil,
		ncnt: 2,
		ecnt: 1,
	}
	g, err := graph.GenerateGraph(sgs)
	if err != want.err {
		t.Errorf("error: got %v, want %v", err, want.err)
	}
	if len(g.GetNodes()) != want.ncnt {
		t.Errorf("node count: got %v, want %v", len(g.GetNodes()), want.ncnt)
	}
	if len(g.GetEdges()) != want.ecnt {
		t.Errorf("edge count: got %v, want %v", len(g.GetEdges()), want.ecnt)
	}
}
