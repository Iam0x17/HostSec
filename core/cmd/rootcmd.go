package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "HostSec",

	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("输出-h查看用法")
	//},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
