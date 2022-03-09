package cmd

import (
	"HostSec/models"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(mkDBCmd)
}

var mkDBCmd = &cobra.Command{
	Use:   "mkdb",
	Short: "compile json into db library",
	Run: func(cmd *cobra.Command, args []string) {
		models.CreateDB()
		models.WriteData2DB()
		defer models.DB.Close()
	},
}
