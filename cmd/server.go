package cmd

import (
	"github.com/qnib/qinventory/server"
	"github.com/spf13/cobra"
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

	/*RootCmd.PersistentFlags().String("collectors", "Gelf,DockerEvents", "Comma separated list of collectors to start")
	viper.BindPFlag("collectors", RootCmd.PersistentFlags().Lookup("collectors"))
	RootCmd.PersistentFlags().Int("ticker-interval", 15000, "Interval of global ticker in milliseconds")
	viper.BindPFlag("ticker-interval", RootCmd.PersistentFlags().Lookup("ticker-interval"))
	RootCmd.PersistentFlags().Int("gelf-port", 12201, "UDP port of GELF collector")
	viper.BindPFlag("gelf-port", RootCmd.PersistentFlags().Lookup("gelf-port"))
	RootCmd.PersistentFlags().String("syslog-output-addr", "127.0.0.1", "UDP target for SyslogOut")
	viper.BindPFlag("syslog-output-addr", RootCmd.PersistentFlags().Lookup("syslog-output-addr"))
	*/
}
