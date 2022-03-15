package attackvector

import (
	"HostSec/attackvector/hips"
	"HostSec/util"
)

func RegOpt(vectorcnname, keyroot, keypath, keyname, keyvalue, opttype string) {
	var res int
	reg := hips.NewRegVector(vectorcnname, keyroot, keypath, keyname, keyvalue)
	switch opttype {
	case "edit":
		res = reg.RegWriteStringValue()
		break
	case "del":
		res = reg.RegDeleteKey()
		break
	}
	util.PrintAttackResult(res, vectorcnname)
}

func CommandOpt(vectorcnname, command string) {
	cmd := hips.NewCommandVector(command)
	util.PrintAttackResult(cmd.ExecCommand(), vectorcnname)
}

func FileOpt(vectorcnname, filepath, filecontent, opttype string) {
	var res int
	file := hips.NewFileVector(filepath, filecontent, opttype)
	switch opttype {
	case "create", "write_append":
		res = file.WriteFile()
		break
	case "del":
		res = file.DelFile()
		break
	}
	util.PrintAttackResult(res, vectorcnname)
}
