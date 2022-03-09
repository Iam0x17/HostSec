package models

import (
	"fmt"
	"github.com/InVisionApp/tabular"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"reflect"
)

var tab tabular.Table
var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("sqlite3", "./attackvector.db")
	if err != nil {
		panic(err)
	}
}

func ShowVectorList() {

	tab = tabular.New()
	tab.Col("v", "攻击向量", 30)
	tab.Col("d", "描述", 30)

	//table := [][]string{}
	vectorListDB := []VectorListDB{}
	DB.Find(&vectorListDB)
	fmt.Println(reflect.TypeOf(vectorListDB))

	format := tab.Print("v", "d")
	for _, v := range vectorListDB {
		fmt.Printf(format, v.VectorName, v.VectorCnName)
	}

}
