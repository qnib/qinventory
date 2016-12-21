package backend

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/gorilla/mux"

	"github.com/mkindahl/gograph/directed"
	qinv "github.com/qnib/qinventory/lib"
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

// GetNode fetches the IDs of all Nodes
func (qg *QGraph) GetNode(id string) (qinv.QEntity, error) {
	if val, ok := qg.Nodes[id]; ok {
		e := val.GetMe()
		return e, nil
	}
	var e qinv.QEntity
	return e, fmt.Errorf("No node with id '%s'", id)
}

// IsSubset returns true if all ids provided are within the Nodes
func (qg *QGraph) IsSubset(keys []string) bool {
	all, _ := qg.GetNodeIDs()
	for i := range all {
		for _, s := range keys {
			if s == all[i] {
				if len(keys) == 1 {
					return true
				}
				keys = append(keys[:i], keys[i+1:]...)
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

// SetNode adds a qinventory entity to the Graph
func (qg QGraph) SetNode(e qinv.QEntity) error {
	nID := e.GetID()
	qg.Nodes[nID] = e.GetMe()
	qg.Graph.AddVertex(nID)
	log.Printf("[II] Added node '%s'\n", nID)
	return nil
}

// AddEdge connects to entities
func (qg *QGraph) AddEdge(x, y qinv.QEntity) (*QGraph, error) {
	ids := []string{x.GetID(), y.GetID()}
	c := []string{x.GetID(), y.GetID()}
	if !qg.IsSubset(ids) {
		msg := fmt.Sprintf("id not in graph: %s, %s", c[0], c[1])
		return qg, errors.New(msg)
	}
	qg.Graph.AddEdge(x.GetID(), y.GetID())
	return qg, nil
}

// StartHttpd creates a http serverPort
func (qg *QGraph) StartHttpd() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", qg.Index).Methods("GET")
	router.HandleFunc("/machines", qg.ListMachines).Methods("GET")
	router.HandleFunc("/machine/{id}", qg.GetMachine).Methods("GET")
	router.HandleFunc("/machine/{id}", qg.AddMachine).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index shows the entry page
func (qg *QGraph) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// ListMachines fetches all machines
func (qg *QGraph) ListMachines(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

// GetMachine fetches one machine
func (qg *QGraph) GetMachine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	m, err := qg.GetNode(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("No node with id '%s'", id), 500)
	}
	b, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// AddMachine adds a machine to the graph
func (qg *QGraph) AddMachine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	decoder := json.NewDecoder(r.Body)
	var m qinv.Machine
	err := decoder.Decode(&m)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	err = qg.SetNode(m.GetMe())
	if err != nil {
		panic(err)
	}
	m.ID = id
	b, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
