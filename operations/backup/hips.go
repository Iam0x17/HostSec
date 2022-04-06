package backup

import (
	"HostSec/attackvector/hips"
	"strings"
)

type BackupRegData struct {
	IsExist  bool
	KeyRoot  string
	KeyPath  string
	KeyName  string
	KeyValue string
}

func FileBackup(filepath, filecontent, opttype string) (int, string) {
	var res int
	var data string
	file := hips.NewFileVector(filepath, filecontent, opttype)
	switch opttype {
	case "create":
		res = 1
		data = ""
		break
	case "write_append", "del":
		res = file.CopyFile(GetDstFile(filepath))
		if res != 0 {
			data = GetDstFile(filepath)
		}
		break
	}
	return res, data
}

func GetDstFile(path string) string {
	str := strings.Split(path, `\`)
	dstFilePath := "./backup/" + str[len(str)-1]
	return dstFilePath
}

func RegBackup(keyroot, keypath, keyname, keyvalue string) (int, string) {
	var backData string
	reg := hips.NewRegistryVector(keyroot, keypath, keyname, keyvalue)
	if reg.RegKeyExists() == 0 {
		backData = "NoKeyPath" + "|" + hips.ReturnExistKeyPath(keyroot, keypath)
		return 1, backData
	}
	res, value := reg.GetKeyValue()
	if res == 0 {
		backData = "NoKeyName" + "|"
		return 1, backData
	}
	if value == "" {
		backData = "NoKeyValue" + "|"
		return 1, backData
	}
	backData = "keyValue" + "|" + value
	return 1, backData
}
