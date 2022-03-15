package control

import (
	"HostSec/attackvector"
	"HostSec/config"
	"HostSec/database"
	"HostSec/util"
	"sync"
)

func regAttackSingle(vectorname string) {
	registerDB := database.RegisterDB{}
	database.DB.Where("vector_name=?", vectorname).Find(&registerDB)
	attackvector.RegOpt(registerDB.VectorCnName, registerDB.KeyRoot, registerDB.KeyPath, registerDB.KeyName, registerDB.KeyValue, registerDB.OptType)
}

func regAttackMulti() {
	var wg sync.WaitGroup
	registerDB := []database.RegisterDB{}
	database.DB.Find(&registerDB)
	//wg.Add(6)
	for _, v := range registerDB {
		wg.Add(1)
		//fmt.Println(v.VectorName)
		go func(v database.RegisterDB) {
			attackvector.RegOpt(v.VectorCnName, v.KeyRoot, v.KeyPath, v.KeyName, v.KeyValue, v.OptType)
			wg.Done()
		}(v)
	}
	wg.Wait()
}

func fileAttackSingle(vectorname string) {
	fileDB := database.FileDB{}
	database.DB.Where("vector_name=?", vectorname).Find(&fileDB)
	attackvector.FileOpt(fileDB.VectorCnName, fileDB.FilePath, fileDB.FileContent, fileDB.OptType)
}

func fileAttackMulti() {
	var wg sync.WaitGroup
	fileDB := []database.FileDB{}
	database.DB.Find(&fileDB)
	//wg.Add(6)
	for _, v := range fileDB {
		wg.Add(1)
		//fmt.Println(v.VectorName)
		go func(v database.FileDB) {
			attackvector.FileOpt(v.VectorCnName, v.FilePath, v.FileContent, v.OptType)
			wg.Done()
		}(v)
	}
	wg.Wait()
}

func commandAttackSingle(vectorname string) {
	commandDB := database.CommandDB{}
	database.DB.Where("vector_name=?", vectorname).Find(&commandDB)
	attackvector.CommandOpt(commandDB.VectorCnName, commandDB.Command)
}

func commandAttackMulti() {
	var wg sync.WaitGroup
	commandDB := []database.CommandDB{}
	database.DB.Find(&commandDB)
	//wg.Add(6)
	for _, v := range commandDB {
		wg.Add(1)
		//fmt.Println(v.VectorName)
		go func(v database.CommandDB) {
			attackvector.CommandOpt(v.VectorCnName, v.Command)
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
		break
	default:
		break
	}

}

func AttackMulti() {

	var wg sync.WaitGroup

	for _, v := range config.AttackDB {
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
		break
	default:
		break
	}
}

func GetAttackType(vectorname string) string {
	vectorListDB := database.VectorListDB{}
	database.DB.Where("vector_name=?", vectorname).Find(&vectorListDB)
	return vectorListDB.Type
}

func Load() {
	util.IsElevated()
	//if util.GetLogSign() {
	//	util.CreateLogFile()
	//}
}

func Unload() {
	database.DB.Close()
	if util.GetLogSign() {
		util.WriteLogResult()
	}
}
