package cmd

import (
	"HostSec/core"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(attackCmd)
	attackCmd.Flags().StringVarP(&vectorName, "vectorname", "n", "", "set attack vector name")
}

var attackCmd = &cobra.Command{
	Use:   "attack",
	Short: "change attack type",
	Run: func(cmd *cobra.Command, args []string) {
		if vectorName == "" {
			core.AttackMulti()
		} else {
			core.AttackSingle(vectorName, core.GetAttackType(vectorName))
		}
	},
}
