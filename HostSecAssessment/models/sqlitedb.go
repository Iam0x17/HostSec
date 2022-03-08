package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("sqlite3", "./attackvector.db")
	if err != nil {
		panic(err)
	}
}

func CreateDB() {
	//db, err := sql.Open("sqlite3", "./attackvector.db")
	////checkErr(err)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//创建表
	sql_table := `
	CREATE TABLE IF NOT EXISTS commanddb(
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    vector_name VARCHAR(64) NULL,
	    vector_cn_name VARCHAR(64) NULL,
	    command VARCHAR(64) NULL
	);
	`
	DB.Exec(sql_table)
	DB.Close()

}

func AddData() {
	//models.CreateDB()
	//registerdb := models.RegisterDB{
	//	VectorName:   "AutoRun",
	//	VectorCnName: "启动项",
	//	KeyRoot:      "HKEY_CURRENT_USER",
	//	KeyPath:      "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run",
	//	KeyName:      "v2rayN",
	//	KeyValue:     "C:\\v2rayN-Core\\v2rayN2.exe",
	//	OptType:      "edit",
	//}
	//err := models.DB.Create(&registerdb).Error

	//filedb := FileDB{
	//	VectorName:   "MalDesktopLNK",
	//	VectorCnName: "恶意创建桌面快捷方式",
	//	FileFullPath: "%Desktop%\\taobao.url",
	//	FileContent:  "[InternetShortcut]\nURL=https://www.taobao.com/",
	//	OptType:      "write",
	//}
	//err := DB.Create(&filedb).Error

	commanddb := CommandDB{
		VectorName:   "CmdAddNetUser",
		VectorCnName: "命令行添加用户账户",
		Command:      "net1 user test test /add",
	}
	err := DB.Create(&commanddb).Error

	if err != nil {
		fmt.Println("---")
		log.Fatal(err)
	}
}
