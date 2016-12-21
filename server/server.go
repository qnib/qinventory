package qserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

// ServeQinventory start daemon
func ServeQinventory(cmd *cobra.Command, args []string) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/machines", ListMachines)
	router.HandleFunc("/machine/{id}", GetMachine)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index shows the entry page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// ListMachines fetches all machines
func ListMachines(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

// GetMachine fetches one machine
func GetMachine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	fmt.Fprintln(w, "Todo show:", ID)
}
