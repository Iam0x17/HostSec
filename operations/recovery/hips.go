package recovery

import (
	"HostSec/attackvector/hips"
	"HostSec/config"
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
		res = file.DelFile()
		break
	}
	//util.PrintAttackResult(res, vectorcnname)
	return res
}
