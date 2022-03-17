package cmd

import (
	"HostSec/controller"
	"HostSec/util"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var vectorName string
var attackLog string

func init() {
	rootCmd.AddCommand(attackCmd)
	attackCmd.Flags().StringVarP(&vectorName, "vectorname", "n", "", "set attack vector name")
	//rootCmd.MarkFlagRequired("vectorname")
	attackCmd.Flags().BoolP("log", "l", false, "enable logging")
	//attackCmd.Flags().StringVar(&attackLog, "attacklog", "l", "output log")
}

var attackCmd = &cobra.Command{
	Use:   "attack",
	Short: "change attack type",
	Run: func(cmd *cobra.Command, args []string) {
		signLog, _ := cmd.Flags().GetBool("log")
		//fmt.Println("logpath " + attackLog)
		util.SetLogSign(signLog)
		if vectorName == "" {
			controller.AttackMulti()
		} else {
			attackType := controller.GetAttackType(vectorName)
			if attackType == "" {
				fmt.Println("攻击向量输入有误，请输出正确的攻击向量")
				ShowVectorList()
				os.Exit(1)
			}
			controller.AttackSingle(vectorName, attackType)
		}
	},
	//Args: cobra.MinimumNArgs(1),
}
