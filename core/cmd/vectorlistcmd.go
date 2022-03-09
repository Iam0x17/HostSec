package cmd

import (
	"HostSec/models"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(vectorListCmd)
}

var vectorListCmd = &cobra.Command{
	Use:   "list",
	Short: "show all attack vector list",
	Run: func(cmd *cobra.Command, args []string) {
		models.ShowVectorList()
	},
}
