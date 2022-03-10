package models

import (
	"fmt"
	"github.com/InVisionApp/tabular"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
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
	tab.Col("o", "序号", 5)
	tab.Col("v", "攻击向量", 30)
	tab.Col("d", "描述", 25)

	//table := [][]string{}
	vectorListDB := []VectorListDB{}
	DB.Find(&vectorListDB)
	//fmt.Println(reflect.TypeOf(vectorListDB))
	nOrder := 1
	format := tab.Print("o", "v", "d")
	for _, v := range vectorListDB {
		fmt.Printf(format, nOrder, v.VectorName, v.VectorCnName)
		nOrder++
	}
}

func FindSingleVector(vectorname string, structdb interface{}) interface{} {
	DB.Where("vector_name=?", vectorname).Find(structdb)
	return structdb
}
