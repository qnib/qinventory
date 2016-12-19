package backend

import (
	"testing"
    "fmt"

	"github.com/mkindahl/gograph/directed"
	qinv "github.com/qnib/qinventory/lib"
	"github.com/stretchr/testify/assert"
)

// Test the instanciation of a PhysicalMachine
func TestNewQGraph(t *testing.T) {
	exp := QGraph{
		Graph: directed.New(),
		Nodes: make(map[string]qinv.QEntity, 0),
	}
	got, _ := NewQGraph()
	assert.Equal(t, exp.Graph, got.Graph)
	assert.EqualValues(t, exp.Nodes, got.Nodes)
}

// Test if add an entity works
func TestAddEntity(t *testing.T) {
	exp := QGraph{
		Graph: directed.New(),
		Nodes: make(map[string]qinv.QEntity, 0),
	}
	m, _ := qinv.NewMachine("host1", "4.7.10")
	exp.Nodes["host1"] = m.GetMe()
	exp.Graph.AddVertex("host1")
	got, _ := NewQGraph()
	got.AddNode(m.GetMe())
	assert.Equal(t, exp.Graph, got.Graph)
	assert.EqualValues(t, exp.Nodes, got.Nodes)

}

// Check IsSubset
func TestIDsAndIsSubset(t *testing.T) {
    g, _ := NewQGraph()
    m1, _ := qinv.NewMachine("host1", "4.7.10")
    m2, _ := qinv.NewMachine("host2", "4.7.10")
    m3, _ := qinv.NewMachine("host3", "4.7.10")
    g.AddNode(m1.GetMe())
    g.AddNode(m2.GetMe())
    g.AddNode(m3.GetMe())
    got, _ := g.GetNodeIDs()
    assert.Equal(t, []string{"host1", "host2", "host3"}, got)
    ids := []string{"host1", "host2"}
    assert.True(t, g.IsSubset(ids))
    ids = []string{"host1", "host4"}
    assert.False(t, g.IsSubset(ids))
}

// Check IDinNodes
func TestIDinNodes(t *testing.T) {
    g, _ := NewQGraph()
    m1, _ := qinv.NewMachine("host1", "4.7.10")
    m2, _ := qinv.NewMachine("host2", "4.7.10")
    m3, _ := qinv.NewMachine("host3", "4.7.10")
    g.AddNode(m1.GetMe())
    g.AddNode(m2.GetMe())
    g.AddNode(m3.GetMe())
    assert.True(t, g.IDinNodes("host1"))
    assert.False(t, g.IDinNodes("host4"))
}

// Check AddEdge
func TestAddEdge(t *testing.T) {
    g, _ := NewQGraph()
    m1, _ := qinv.NewMachine("host1", "4.7.10")
    m2, _ := qinv.NewMachine("host2", "4.7.10")
    m3, _ := qinv.NewMachine("host3", "4.7.10")
    g.AddNode(m1.GetMe())
    g.AddNode(m2.GetMe())
    g.AddEdge(m1.GetMe(), m2.GetMe())
    assert.True(t, g.Graph.HasEdge(m1.GetID(), m2.GetID()))
    assert.False(t, g.Graph.HasEdge(m2.GetID(), m1.GetID()))
    _, err := g.AddEdge(m1.GetMe(), m3.GetMe())
    assert.Equal(t, "id not in graph: host1, host3", fmt.Sprintf("%s", err))

}
