package main

import "testing"

func TestCreateSgNode(t *testing.T) {
	sg := &SecurityGroup{
		GroupID:     "id",
		GroupName:   "name",
		Description: "description",
	}
	got := CreateSgNode(sg)
	want := &Node{
		id:   "id",
		text: "id:name\n(description)",
	}
	if *got != *want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCreateIPNode(t *testing.T) {
	ip := &IPRange{
		CidrIP:      "cidr",
		Description: "description",
	}
	got := CreateIPNode(ip)
	want := &Node{
		id:   "cidr",
		text: "cidr\n(description)",
	}
	if *got != *want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCreateUserIDGroupPairNode(t *testing.T) {
	userIDGroupPair := &UserIDGroupPair{
		GroupID: "id",
		UserID:  "user",
	}
	got := CreateUserIDGroupPairNode(userIDGroupPair)
	want := &Node{
		id:   "id",
		text: "id\n(user)",
	}
	if *got != *want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCreatePrefixNode(t *testing.T) {
	prefix := &PrefixListId{
		PrefixListID: "id",
		Description:  "description",
	}
	got := CreatePrefixNode(prefix)
	want := &Node{
		id:   "id",
		text: "id\n(description)",
	}
	if *got != *want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCreateIpv6Node(t *testing.T) {
	ipv6 := &Ipv6Range{
		CidrIpv6:    "cidr",
		Description: "description",
	}
	got := CreateIpv6Node(ipv6)
	want := &Node{
		id:   "cidr",
		text: "cidr\n(description)",
	}
	if *got != *want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCreatePermEdge(t *testing.T) {
	got := CreatePermEdge("from", "to", 1, 2)
	want := &Edge{
		from: "from",
		to:   "to",
		text: "1-2",
	}
	if *got != *want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestAddNode(t *testing.T) {
	g := NewGraph()
	n := &Node{id: "id"}
	g.AddNode(n)
	got := g.nodes["id"]
	if got != n {
		t.Errorf("got %v, want %v", got, n)
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph()
	e := &Edge{}
	g.AddEdge(e)
	got := g.edges[0]
	if got != e {
		t.Errorf("got %v, want %v", got, e)
	}
}
