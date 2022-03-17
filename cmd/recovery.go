package cmd

import (
	"HostSec/control"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(recoveryCmd)
}

var recoveryCmd = &cobra.Command{
	Use:   "recovery",
	Short: "recovery attack vector",
	Run: func(cmd *cobra.Command, args []string) {
		control.RecoveryEnv()
	},
}
