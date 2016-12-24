package cmd

import (
	"github.com/qnib/qinventory/server"
	"github.com/spf13/cobra"
    "github.com/spf13/viper"

)

// watchSrv loops over nodes, services and tasks
var qServer = &cobra.Command{
	Use:   "server",
	Short: "Starts server daemon",
	Long:  ``,
	Run:   qserver.ServeQinventory,
}

func init() {
	RootCmd.AddCommand(qServer)
	RootCmd.PersistentFlags().Bool("enable-zipkin", false, "Start http-server with zipkin tracing")
	viper.BindPFlag("enable-zipkin", RootCmd.PersistentFlags().Lookup("enable-zipkin"))
	RootCmd.PersistentFlags().String(
		"zipkin-url",
		"http://localhost:9411/api/v1/spans",
		"URL to zipkin server",
	)
}
