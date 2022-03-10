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
		//models.WriteData2DB()
		models.WriteData2DBSingle()
		defer models.DB.Close()
	},
}
