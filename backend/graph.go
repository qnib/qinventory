package backend

import (
	"github.com/mkindahl/gograph/directed"
	qinv "github.com/qnib/qinventory/lib"
    "sort"
)

// QGraph holds the inventory graph
type QGraph struct {
	Graph *directed.Graph
	Nodes map[string]qinv.QEntity
}

// NewQGraph returns a new Graph
func NewQGraph() (QGraph, error) {
	return QGraph{
		Graph: directed.New(),
		Nodes: make(map[string]qinv.QEntity, 0),
	}, nil
}

// GetNodeIDs fetches the IDs of all Nodes
func (qg *QGraph) GetNodeIDs() ([]string, error) {
    var ids []string
    for k := range qg.Nodes {
        ids = append(ids, k)
    }
    sort.Strings(ids)
    return ids, nil
}

// IsSubset returns true if all ids provided are within the Nodes
func (qg *QGraph) IsSubset(ids []string) bool {
    all, _ := qg.GetNodeIDs()
    for i := range all {
        for _, s := range ids {
            if s == all[i] {
                if len(ids) == 1 {
                    return true
                }
                ids = append(ids[:i], ids[i+1:]...)
            }
        }
    }
    return false
}

// IDinNodes returns true if id is part of nodes
func (qg *QGraph) IDinNodes(id string) bool {
    all, _ := qg.GetNodeIDs()
    for i := range all {
        if id == all[i] {
            return true
        }
    }
    return false
}

// AddNode adds a qinventory entity to the Graph
func (qg *QGraph) AddNode(e qinv.QEntity) (*QGraph, error) {
	nID := e.GetID()
	qg.Nodes[nID] = e.GetMe()
	qg.Graph.AddVertex(nID)
	return qg, nil
}

/*
// AddEdge connects to entities
func (qg *QGraph) AddEdge(x,y qinv.QEntity) (*QGraph, error) {
    for k := range qg.Nodes {

    }
    qg.Graph.AddEdge(x.GetID(), y.GetID())
    return qg, nil
}
*/
