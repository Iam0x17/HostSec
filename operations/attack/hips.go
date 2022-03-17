package attack

import (
	"HostSec/attackvector/hips"
)

func RegistryAttack(vectorcnname, keyroot, keypath, keyname, keyvalue, opttype string) int {
	var res int
	reg := hips.NewRegistryVector(vectorcnname, keyroot, keypath, keyname, keyvalue)
	switch opttype {
	case "edit":
		res = reg.RegWriteStringValue()
		break
	case "del":
		res = reg.RegDeleteKey()
		break
	}
	//util.PrintAttackResult(res, vectorcnname)
	return res
}

func CommandAttack(vectorcnname, command string) int {
	cmd := hips.NewCommandVector(command)
	res := cmd.ExecCommand()
	//util.PrintAttackResult(res, vectorcnname)
	return res
}

func FileAttack(vectorcnname, filepath, filecontent, opttype string) int {
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
	//util.PrintAttackResult(res, vectorcnname)
	return res
}
