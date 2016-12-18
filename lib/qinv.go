package qinv

import (
    //"github.com/twmb/algoimpl/go/graph"
    //"encoding/json"
)

// QInventory provides an interface for all backends
type QInventory interface {
    AddMachine(m Machine)
    AddContainer(c Container)
}

/*
// QGraph holds the inventory graph
type QGraph struct {
    Graph *graph.Graph
    Nodes map[string]graph.Node
}

// NewQGraph returns an instance
func NewQGraph() (QGraph, error) {
    return QGraph{
        Graph: graph.New(graph.Directed),
        Nodes: make(map[string]graph.Node, 0),
    }, nil
}

// AddMachine adds a machine to the Graph
func (qg *QGraph) AddMachine(m Machine) *QGraph {
    qg.Nodes[m.Hostname] = qg.Graph.MakeNode()
    b, _ := json.Marshal(m)
    *qg.Nodes[m.Hostname].Value = b
    return qg
}
*/
