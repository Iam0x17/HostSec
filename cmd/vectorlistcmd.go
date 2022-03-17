package cmd

import (
	"HostSec/database"
	"fmt"
	"github.com/InVisionApp/tabular"
	"github.com/spf13/cobra"
)

//var tab tabular.Table

func init() {
	rootCmd.AddCommand(vectorListCmd)
}

var vectorListCmd = &cobra.Command{
	Use:   "list",
	Short: "show all attack vector list",
	Run: func(cmd *cobra.Command, args []string) {
		ShowVectorList()
	},
}

func ShowVectorList() {
	var tab tabular.Table
	tab = tabular.New()
	tab.Col("o", "序号", 5)
	tab.Col("v", "攻击向量", 30)
	tab.Col("d", "描述", 25)

	//table := [][]string{}
	vectorListDB := []database.VectorListDB{}
	database.DB.Find(&vectorListDB)
	//fmt.Println(reflect.TypeOf(vectorListDB))
	nOrder := 1
	format := tab.Print("o", "v", "d")
	for _, v := range vectorListDB {
		fmt.Printf(format, nOrder, v.VectorName, v.VectorCnName)
		nOrder++
	}
}
