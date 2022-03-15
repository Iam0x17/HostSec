package cmd

import (
	"HostSec/database"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(mkDBCmd)
}

var mkDBCmd = &cobra.Command{
	Use:   "mkdb",
	Short: "compile json into db library",
	Run: func(cmd *cobra.Command, args []string) {
		database.CreateDB()
		//models.WriteData2DB()
		database.WriteData2DBSingle()
		defer database.DB.Close()
	},
}
