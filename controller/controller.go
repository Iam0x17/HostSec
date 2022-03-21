package controller

import (
	"HostSec/config"
	"HostSec/database"
	"HostSec/operations/attack"
	"HostSec/operations/backup"
	"HostSec/operations/log"
	"HostSec/operations/recovery"
	"HostSec/util"
	"fmt"
	"sync"
)

var record = log.NewRecord()

type zc log.RecordData

func RecoveryEnv() {
	recordData := []log.RecordData{}
	database.RecoveryDB.Find(&recordData)

	for _, v := range recordData {
		RecoveryRecord(v)
	}
}

func RecoveryRecord(recorddata log.RecordData) {
	var res int
	switch recorddata.AttackType {
	case "register":
		res = recovery.RegRecovery(recorddata.RawData, recorddata.BackupData)
		break
	case "file":
		res = recovery.FileRecovery(recorddata.RawData, recorddata.BackupData, recorddata.RecoveryType)
		break
	case "command":
		res = recovery.CmdRecovery(recorddata.BackupData)
	default:
		break
	}

	if res != 0 {
		rSuccess := recorddata.VectorCnName + " 恢复成功"
		fmt.Println(rSuccess)
		database.RecoveryDB.Delete(&recorddata)
	}
}

func AddRecordData2DB(vectorcnname, attacktype, rawdata, backupdata, optType string) {
	rData := log.RecordData{
		AttackType:   attacktype,
		VectorCnName: vectorcnname,
		RawData:      rawdata,
		BackupData:   backupdata,
		RecoveryType: recovery.GetRecoveryType(attacktype, optType),
	}
	database.RecoveryDB.Create(&rData)
	//record.Set(vectorcnname, &rData)
}

func AttackRecord(vectorcnname, attacktype string, dbtype interface{}) {
	var resAttack int
	var resBackup int
	var dataBackup string
	var optType string
	var dataRaw string

	switch attacktype {
	case "register":
		v := dbtype.(database.RegisterDB)
		resBackup, dataBackup = backup.RegBackup(v.KeyRoot, v.KeyPath, v.KeyName, v.KeyValue)
		if resBackup == 1 {
			resAttack = attack.RegistryAttack(v.KeyRoot, v.KeyPath, v.KeyName, v.KeyValue, v.OptType)
		}
		optType = v.OptType
		dataRaw = v.KeyRoot + "|" + v.KeyPath + "|" + v.KeyName + "|" + v.KeyValue
	case "file":
		v := dbtype.(database.FileDB)
		resBackup, dataBackup = backup.FileBackup(v.FilePath, v.FileContent, v.OptType)
		if resBackup == 1 {
			resAttack = attack.FileAttack(v.FilePath, v.FileContent, v.OptType)
		}
		optType = v.OptType
		dataRaw = v.FilePath
		break
	case "command":
		v := dbtype.(database.CommandDB)
		resBackup = v.Backup
		resAttack = attack.CommandAttack(v.Command)
		if resBackup != 0 {
			dataBackup = v.Recovery
		}
		break
	}

	if resBackup != 0 && resAttack != 0 {
		AddRecordData2DB(vectorcnname, attacktype, dataRaw, dataBackup, optType)
	}

	util.PrintAttackResult(resAttack, vectorcnname)
}

func regAttackSingle(vectorname string) {
	registerDB := database.RegisterDB{}
	database.DB.Where("vector_name=?", vectorname).Find(&registerDB)
	//attack.RegistryOpt(registerDB.VectorCnName, registerDB.KeyRoot, registerDB.KeyPath, registerDB.KeyName, registerDB.KeyValue, registerDB.OptType)
	AttackRecord(registerDB.VectorCnName, "register", registerDB)
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
			//attack.RegistryOpt(v.VectorCnName, v.KeyRoot, v.KeyPath, v.KeyName, v.KeyValue, v.OptType)
			AttackRecord(v.VectorCnName, "register", v)
			wg.Done()
		}(v)
	}
	wg.Wait()
}

func fileAttackSingle(vectorname string) {
	fileDB := database.FileDB{}
	database.DB.Where("vector_name=?", vectorname).Find(&fileDB)
	//attack.FileOpt(fileDB.VectorCnName, fileDB.FilePath, fileDB.FileContent, fileDB.OptType)
	AttackRecord(fileDB.VectorCnName, "file", fileDB)
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
			//attack.FileOpt(v.VectorCnName, v.FilePath, v.FileContent, v.OptType)
			AttackRecord(v.VectorCnName, "file", v)
			wg.Done()
		}(v)
	}
	wg.Wait()
}

func commandAttackSingle(vectorname string) {
	commandDB := database.CommandDB{}
	database.DB.Where("vector_name=?", vectorname).Find(&commandDB)
	//attack.CommandOpt(commandDB.VectorCnName, commandDB.Command)
	AttackRecord(commandDB.VectorCnName, "command", commandDB)
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
			//attack.CommandOpt(v.VectorCnName, v.Command)
			AttackRecord(v.VectorCnName, "command", v)
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
