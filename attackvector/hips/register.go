package hips

import (
	registry "golang.org/x/sys/windows/registry"
	"strings"
	"syscall"
)

type RegistrySrv interface {
	RegWriteStringValue(keyvalue ...string) int
	RegDeleteKey() int
	RegKeyExists() int
	GetKeyValue() (int, string)
	RegDeleteKeyValue() int
}

type RegistryData struct {
	KeyRoot  registry.Key
	KeyPath  string
	KeyName  string
	KeyValue string
}

func NewRegistryVector(keyroot, keypath, keyname, keyvalue string) RegistrySrv {
	regData := RegistryData{
		KeyRoot:  getKeyRoot(keyroot),
		KeyPath:  keypath,
		KeyName:  keyname,
		KeyValue: keyvalue,
	}

	return regData
}

func (reg RegistryData) RegKeyExists() int {

	var result int

	k, err := registry.OpenKey(reg.KeyRoot, reg.KeyPath, registry.QUERY_VALUE)
	if err != nil {
		//log.Fatal(err)
		result = 0
	} else {
		result = 1
	}

	defer k.Close()

	return result
}

func (reg RegistryData) GetKeyValue() (int, string) {
	k, errOpenKey := registry.OpenKey(reg.KeyRoot, reg.KeyPath, registry.QUERY_VALUE)
	if errOpenKey != nil {
		return -1, ""
	}
	defer k.Close()
	value, _, errGetValue := k.GetStringValue(reg.KeyName)
	if errGetValue != nil {
		//if int(errGetValue.(syscall.Errno)) == 2 {
		//	return 0, ""
		//}
		//return -1, ""
		return 0, ""
	}
	return 1, value
}

//获取原始注册表存在最初的键
func ReturnExistKeyPath(keyroot, keypath string) string {

	k, err := registry.OpenKey(getKeyRoot(keyroot), keypath, registry.QUERY_VALUE)
	//k.GetMUIStringValue()
	if err != nil {
		if int(err.(syscall.Errno)) == 2 {
			index := strings.LastIndex(keypath, "\\")
			if index == -1 {
				return ""
			}
			ReturnExistKeyPath(keyroot, keypath[:index])
		}

	}
	defer k.Close()

	return keypath
}

func (reg RegistryData) RegWriteKey() {

	k, _, _ := registry.CreateKey(reg.KeyRoot, reg.KeyPath, registry.ALL_ACCESS)
	defer k.Close()
}

//设置注册表键值
func (reg RegistryData) RegWriteStringValue(keyvalue ...string) int {
	var result int
	var errSetValue error
	k, _, errhandle := registry.CreateKey(reg.KeyRoot, reg.KeyPath, registry.ALL_ACCESS)
	if errhandle != nil {
		return 2
	}
	defer k.Close()
	if len(keyvalue) == 0 {
		errSetValue = k.SetStringValue(reg.KeyName, reg.KeyValue)
	} else {
		errSetValue = k.SetStringValue(reg.KeyName, keyvalue[0])
	}

	if errSetValue != nil {
		result = 0
	} else {
		result = 1
	}

	return result
}

func (reg RegistryData) RegDeleteKey() int {
	var result int

	err := registry.DeleteKey(reg.KeyRoot, reg.KeyPath)
	if err != nil {
		result = 0
	} else {
		result = 1
	}

	return result
}

func RegDeleteMulKey(keyroot, keypath string) int {

	err := registry.DeleteKey(getKeyRoot(keyroot), keypath)
	if err != nil {
		return 0
	} else {
		tmpKeyPath := ReturnExistKeyPath(keyroot, keypath)
		if tmpKeyPath != keypath {
			RegDeleteMulKey(keyroot, tmpKeyPath)
		}
	}

	return 1
}

func (reg RegistryData) RegDeleteKeyValue() int {

	k, errOpenKey := registry.OpenKey(reg.KeyRoot, reg.KeyPath, registry.ALL_ACCESS)
	if errOpenKey != nil {
		return 0
	}
	defer k.Close()

	errDelValue := k.DeleteValue(reg.KeyName)
	if errDelValue != nil {
		return 0
	}

	return 1
}

func getKeyRoot(keyroot string) registry.Key {

	var keyRoot registry.Key

	switch keyroot {
	case "HKEY_CLASSES_ROOT":
		keyRoot = registry.CLASSES_ROOT
		break
	case "HKEY_CURRENT_USER":
		keyRoot = registry.CURRENT_USER
		break
	case "HKEY_LOCAL_MACHINE":
		keyRoot = registry.LOCAL_MACHINE
		break
	case "HKEY_USERS":
		keyRoot = registry.USERS
		break
	case "HKEY_CURRENT_CONFIG":
		keyRoot = registry.CURRENT_CONFIG
		break
	default:
	}

	return keyRoot
}
