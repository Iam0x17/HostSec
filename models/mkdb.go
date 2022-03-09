package models

import (
	"HostSec/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

func CreateDB() {
	//创建表
	sql_registerdb_table := `
	CREATE TABLE IF NOT EXISTS registerdb(
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    vector_name VARCHAR(64) NULL,
	    vector_cn_name VARCHAR(64) NULL,
		key_root VARCHAR(64) NULL,
		key_path VARCHAR(64) NULL,
		key_name VARCHAR(64) NULL,
		key_value VARCHAR(64) NULL,
		opt_type VARCHAR(64) NULL
	);
	`

	sql_filedb_table := `
	CREATE TABLE IF NOT EXISTS filedb(
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    vector_name VARCHAR(64) NULL,
	    vector_cn_name VARCHAR(64) NULL,
	    file_full_path VARCHAR(64) NULL,
		file_content VARCHAR(64) NULL,
	    opt_type VARCHAR(64) NULL
	);
	`

	sql_commanddb_table := `
	CREATE TABLE IF NOT EXISTS commanddb(
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    vector_name VARCHAR(64) NULL,
	    vector_cn_name VARCHAR(64) NULL,
	    command VARCHAR(64) NULL
	);
	`

	sql_vectorlistdb_table := `
	CREATE TABLE IF NOT EXISTS vectorlistdb(
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    vector_name VARCHAR(64) NULL,
	    vector_cn_name VARCHAR(64) NULL,
	    type VARCHAR(64) NULL
	);
	`

	DB.Exec(sql_registerdb_table)
	DB.Exec(sql_filedb_table)
	DB.Exec(sql_commanddb_table)
	DB.Exec(sql_vectorlistdb_table)

}

func writeFileData2DB() {
	DB.Exec("DELETE FROM filedb")
	bytes, err := ioutil.ReadFile("./attackvector/file.json")
	if err != nil {
		fmt.Println("读取register.json文件失败", err)
		return
	}

	fileDB := []FileDB{}
	err1 := json.Unmarshal(bytes, &fileDB)
	if err1 != nil {
		fmt.Println("err = ", err1)
		return
	}
	for _, v := range fileDB {
		errWriteDB := DB.Create(&v).Error
		if errWriteDB != nil {
			log.Fatal(errWriteDB)
		}
		//fmt.Printf("tmp = %+v\n", v.VectorName)
	}
}

func writeRegisterData2DB() {
	DB.Exec("DELETE FROM registerdb")
	bytes, err := ioutil.ReadFile("./attackvector/register.json")
	if err != nil {
		fmt.Println("读取register.json文件失败", err)
		return
	}

	registerDB := []RegisterDB{}
	err1 := json.Unmarshal(bytes, &registerDB)
	if err1 != nil {
		fmt.Println("err = ", err1)
		return
	}
	for _, v := range registerDB {
		errWriteDB := DB.Create(&v).Error
		if errWriteDB != nil {
			log.Fatal(errWriteDB)
		}
		//fmt.Printf("tmp = %+v\n", v.VectorName)
	}
}

func writeCommandData2DB() {
	DB.Exec("DELETE FROM commanddb")
	bytes, err := ioutil.ReadFile("./attackvector/command.json")
	if err != nil {
		fmt.Println("读取command.json文件失败", err)
		return
	}

	commandDB := []CommandDB{}
	err1 := json.Unmarshal(bytes, &commandDB)
	if err1 != nil {
		fmt.Println("err = ", err1)
		return
	}
	for _, v := range commandDB {
		errWriteDB := DB.Create(&v).Error
		if errWriteDB != nil {
			log.Fatal(errWriteDB)
		}
		//fmt.Printf("tmp = %+v\n", v.VectorName)
	}
}

func GetJsonData(filepath string) {

	var errJson error
	vectorListDB := VectorListDB{}
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("读取文件失败", err)
		return
	}

	attackSilce := []string{"file", "command", "register"}
	isExists, jsonType := util.ExistWhoStr(filepath, attackSilce)
	if isExists == false {
		fmt.Println("缺少必要的json文件")
		os.Exit(1)
	}
	switch jsonType {
	case "file":
		fileDB := []FileDB{}
		errJson = json.Unmarshal(bytes, &fileDB)
		if errJson != nil {
			fmt.Println("err = ", errJson)
			return
		}
		for _, v := range fileDB {
			vectorListDB.VectorName = v.VectorName
			vectorListDB.VectorCnName = v.VectorCnName
			vectorListDB.Type = jsonType
			errWriteDB := DB.Create(vectorListDB).Error
			if errWriteDB != nil {
				log.Fatal(errWriteDB)
			}
		}
		break
	case "command":
		commandDB := []CommandDB{}
		errJson = json.Unmarshal(bytes, &commandDB)
		if errJson != nil {
			fmt.Println("err = ", errJson)
			return
		}
		for _, v := range commandDB {
			vectorListDB.VectorName = v.VectorName
			vectorListDB.VectorCnName = v.VectorCnName
			vectorListDB.Type = jsonType
			errWriteDB := DB.Create(vectorListDB).Error
			if errWriteDB != nil {
				log.Fatal(errWriteDB)
			}
		}
		break
	case "register":
		registerDB := []RegisterDB{}
		errJson = json.Unmarshal(bytes, &registerDB)
		if errJson != nil {
			fmt.Println("err = ", errJson)
			return
		}
		for _, v := range registerDB {
			vectorListDB.VectorName = v.VectorName
			vectorListDB.VectorCnName = v.VectorCnName
			vectorListDB.Type = jsonType
			errWriteDB := DB.Create(vectorListDB).Error
			if errWriteDB != nil {
				log.Fatal(errWriteDB)
			}
		}
		break
	}

}

func writeVectorListData2DB() {
	DB.Exec("DELETE FROM vectorlistdb")
	fileSlice := util.GetFileList("./attackvector")
	for _, v := range fileSlice {
		GetJsonData(v)
	}
}

func WriteData2DB() {
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		writeFileData2DB()
		wg.Done()
	}()
	go func() {
		writeRegisterData2DB()
		wg.Done()
	}()
	go func() {
		writeCommandData2DB()
		wg.Done()
	}()
	go func() {
		writeVectorListData2DB()
		wg.Done()
	}()
	wg.Wait()
}
