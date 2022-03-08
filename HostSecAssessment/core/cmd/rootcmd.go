package cmd

import (
	"HostSecAssessment/core"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var vectorName string

func init() {
	AddVectorListCmd(vectorListCmd)
	AddAttackCmd(attackCmd)
}

var rootCmd = &cobra.Command{
	Use: "HostSecAssessment",

	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println(args)
	//},
}

func AddVectorListCmd(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func AddAttackCmd(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
	attackCmd.Flags().StringVarP(&vectorName, "vectorname", "n", "", "set attack vector name")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var vectorListCmd = &cobra.Command{
	Use:   "list",
	Short: "show all attack vector list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deal")
	},
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
