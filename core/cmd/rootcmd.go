package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var vectorName string

var rootCmd = &cobra.Command{
	Use: "HostSec",

	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println(args)
	//},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
