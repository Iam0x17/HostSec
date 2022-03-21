package recovery

import (
	"HostSec/attackvector/hips"
	"HostSec/config"
	"strings"
)

func GetRecoveryType(attacktype, opttype string) string {
	switch attacktype {
	case config.HipsFile:
		if opttype == "create" {
			return "del"
		} else if opttype == "write_append" {
			return "del_copy"
		} else if opttype == "del" {
			return "copy"
		}
		break
	case config.HipsReg:
		if opttype == "edit" {
			return "edit"
		} else {
			return "add"
		}
	}
	return ""
}

func FileRecovery(rawpath, backuppath, recoverytype string) int {
	var res int
	file := hips.NewFileVector(backuppath, "", recoverytype)
	switch recoverytype {
	case "del_copy":
		res = file.DelCopy(rawpath)
		if res != 0 {
			file.DelFile()
		}
		break
	case "del":
		res = file.DelFile(rawpath)
		break
	case "copy":
		res = file.CopyFile(rawpath)
	}
	//util.PrintAttackResult(res, vectorcnname)
	return res
}

func RegRecovery(rawreg, backupreg string) int {
	var res int
	rawData := strings.Split(rawreg, "|")
	backupData := strings.Split(backupreg, "|")
	reg := hips.NewRegistryVector(rawData[0], rawData[1], rawData[2], rawData[3])

	switch backupData[0] {
	case "keyValue":
		res = reg.RegWriteStringValue(backupData[1])
		break
	case "NoKeyPath":
		res = hips.RegDeleteMulKey(rawData[0], rawData[1])
		break
	case "NoKeyName":
		res = reg.RegDeleteKeyValue()
		break
	case "NokeyValue":
		res = reg.RegWriteStringValue("")
		break
	}

	//util.PrintAttackResult(res, vectorcnname)
	return res
}

func CmdRecovery(backupcmd string) int {
	cmd := hips.NewCommandVector(backupcmd)
	res := cmd.ExecCommand()
	//util.PrintAttackResult(res, vectorcnname)
	return res
}
