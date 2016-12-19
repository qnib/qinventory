package backend

import (
	"github.com/mkindahl/gograph/directed"
	qinv "github.com/qnib/qinventory/lib"
	//"github.com/qnib/qinventory/lib/"
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

// AddNode adds a qinventory entity to the Graph
func (qg *QGraph) AddNode(e qinv.QEntity) (*QGraph, error) {
	nID := e.GetID()
	qg.Nodes[nID] = e.GetMe()
	qg.Graph.AddVertex(nID)
	return qg, nil
}
