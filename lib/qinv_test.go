package qinv

/*
import (
    "testing"
    "fmt"

    "github.com/stretchr/testify/assert"
    //"github.com/twmb/algoimpl/go/graph"
)

// Test the instanciation of a PhysicalMachine
func TestNewQGraph(t *testing.T) {
    g := graph.New(graph.Directed)
    exp := QGraph{
        Graph: g,
        Nodes: make(map[string]graph.Node, 0),
    }
    got, _ := NewQGraph()
    assert.Equal(t, exp, got)
}

// TestAddMachine adds a machine to the Graph
func TestAddMachine(t *testing.T) {
    m := Machine{
        Hostname: "host1",
        Kernel: "4.7.10",
    }

    g := graph.New(graph.Directed)
    // Add node to graph
    qg, _ := NewQGraph()
    qg.AddMachine(m)
    // Expect Graph with one node
    n := g.MakeNode()
    *n.Value = map[string]string{
        "hostname": "host1",
        "kernel": "4.7.10",
    }
    nodes := make(map[string]graph.Node, 0)
    nodes["host1"] = n
    expG := QGraph{
        Graph: g,
        Nodes: nodes,
    }
    got := qg.Nodes["host1"]
    exp := expG.Nodes["host1"]
    assert.Equal(t, exp.Value, got.Value, fmt.Sprintf("%v", exp.Value["hostname"]))
}
*/
