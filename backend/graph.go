package backend

import (
	"github.com/mkindahl/gograph/directed"
)

// QGraph holds the inventory graph
type QGraph struct {
	Graph *directed.Graph
}

// NewQGraph returns a new Graph
func NewQGraph() (QGraph, error) {
	return QGraph{
		Graph: directed.New(),
	}, nil
}
