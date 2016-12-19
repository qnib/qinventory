package backend

import (
	"testing"

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
