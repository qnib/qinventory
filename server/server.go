package qserver

import (
	"github.com/qnib/qinventory/backend"
	qinv "github.com/qnib/qinventory/lib"

	"github.com/spf13/cobra"
	"github.com/spf13/cast"
)

// ServeQinventory start daemon
func ServeQinventory(cmd *cobra.Command, args []string) {
	qg, _ := backend.NewQGraph()
	m, _ := qinv.NewMachine("host1", "4.7.10")
	qg.AddNode(m.GetMe())
	cfg := map[string]string{
		"enable-zipkin": cast.ToString(cmd.Flag("enable-zipkin").Value),
		"enable-url": cast.ToString(cmd.Flag("enable-url").Value),
	}
	qg.StartHttpd(cfg)
}
