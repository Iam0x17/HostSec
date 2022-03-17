package hips

import (
	registry "golang.org/x/sys/windows/registry"
)

type RegistrySrv interface {
	RegWriteStringValue() int
	RegDeleteKey() int
}

type RegistryData struct {
	VectorName   string
	VectorCnName string
	KeyRoot      registry.Key
	KeyPath      string
	KeyName      string
	KeyValue     string
}

func NewRegistryVector(vectorcnname, keyroot, keypath, keyname, keyvalue string) RegistrySrv {
	regData := RegistryData{
		VectorCnName: vectorcnname,
		KeyRoot:      getKeyRoot(keyroot),
		KeyPath:      keypath,
		KeyName:      keyname,
		KeyValue:     keyvalue,
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

func (reg RegistryData) RegWriteKey() {

	k, _, _ := registry.CreateKey(reg.KeyRoot, reg.KeyPath, registry.ALL_ACCESS)
	defer k.Close()
}

//设置注册表键值
func (reg RegistryData) RegWriteStringValue() int {
	var result int

	k, _, errhandle := registry.CreateKey(reg.KeyRoot, reg.KeyPath, registry.ALL_ACCESS)
	if errhandle != nil {
		return 2
	}
	defer k.Close()
	err := k.SetStringValue(reg.KeyName, reg.KeyValue)
	if err != nil {
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
