package core

import (
	"HostSecAssessment/models"
	"HostSecAssessment/operations"
	"HostSecAssessment/util"
	"sync"
)

func regAttackSingle(vectorname string) {
	registerDB := models.RegisterDB{}
	models.DB.Where("vector_name=?", vectorname).Find(&registerDB)
	operations.RegOpt(registerDB.VectorCnName, registerDB.KeyRoot, registerDB.KeyPath, registerDB.KeyName, registerDB.KeyValue, registerDB.OptType)
}

func regAttackMulti() {
	var wg sync.WaitGroup
	registerDB := []models.RegisterDB{}
	models.DB.Find(&registerDB)
	//wg.Add(6)
	for _, v := range registerDB {
		wg.Add(1)
		//fmt.Println(v.VectorName)
		go func(v models.RegisterDB) {
			operations.RegOpt(v.VectorCnName, v.KeyRoot, v.KeyPath, v.KeyName, v.KeyValue, v.OptType)
			wg.Done()
		}(v)
	}
	wg.Wait()
}

func fileAttackSingle(vectorname string) {
	fileDB := models.FileDB{}
	models.DB.Where("vector_name=?", vectorname).Find(&fileDB)
	operations.FileOpt(fileDB.VectorCnName, fileDB.FileFullPath, fileDB.FileContent, fileDB.OptType)
}

func fileAttackMulti() {
	var wg sync.WaitGroup
	fileDB := []models.FileDB{}
	models.DB.Find(&fileDB)
	//wg.Add(6)
	for _, v := range fileDB {
		wg.Add(1)
		//fmt.Println(v.VectorName)
		go func(v models.FileDB) {
			operations.FileOpt(v.VectorCnName, v.FileFullPath, v.FileContent, v.OptType)
			wg.Done()
		}(v)
	}
	wg.Wait()
}

func commandAttackSingle(vectorname string) {
	commandDB := models.CommandDB{}
	models.DB.Where("vector_name=?", vectorname).Find(&commandDB)
	operations.CcommandOpt(commandDB.VectorCnName, commandDB.Command)
}

func commandAttackMulti() {
	var wg sync.WaitGroup
	commandDB := []models.CommandDB{}
	models.DB.Find(&commandDB)
	//wg.Add(6)
	for _, v := range commandDB {
		wg.Add(1)
		//fmt.Println(v.VectorName)
		go func(v models.CommandDB) {
			operations.CcommandOpt(v.VectorCnName, v.Command)
			wg.Done()
		}(v)
	}
	wg.Wait()
}

func attackMultiType(attacktype string) {
	switch attacktype {
	case "register":
		regAttackMulti()
		break
	case "file":
		fileAttackMulti()
		break
	case "command":
		commandAttackMulti()
	}

}

func AttackMulti() {

	var wg sync.WaitGroup
	attack := []string{"register", "file", "command"}

	for _, v := range attack {
		wg.Add(1)
		//fmt.Println(v.VectorName)
		go func(v string) {
			attackMultiType(v)
			wg.Done()
		}(v)
	}
	wg.Wait()

}

func AttackSingle(vectorname, attacktype string) {
	switch attacktype {
	case "register":
		regAttackSingle(vectorname)
		break
	case "file":
		fileAttackSingle(vectorname)
		break
	case "command":
		commandAttackSingle(vectorname)
	}
}

func GetAttackType(vectorname string) string {
	return "file"
}

func Unload() {
	models.DB.Close()
	util.WriteLogResult()
}
