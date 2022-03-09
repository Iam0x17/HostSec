package operations

import (
	"HostSec/util"
	registry "golang.org/x/sys/windows/registry"
)

func RegOpt(vectorcnname, keyroot, keypath, keyname, keyvalue, opttype string) {
	//wg *sync.WaitGroup,
	//IsExist := regKeyExists(keyroot, keypath)
	//if IsExist != 1 {
	//	//RegWriteKey(keyroot, keypath)
	//	//log.Fatal("注册表不存在")
	//	fmt.Println("注册表键不存在")
	//}
	var res int

	switch opttype {
	case "edit":
		res = regWriteStringValue(keyroot, keypath, keyname, keyvalue)
		break
	case "del":
		res = regDeleteKey(keyroot, keypath)
		break
	}
	util.PrintAttackResult(res, vectorcnname)
}

func getHKEY(root string) registry.Key {

	var HKEY registry.Key

	if root == "HKEY_CLASSES_ROOT" {
		HKEY = registry.CLASSES_ROOT
	} else if root == "HKEY_CURRENT_USER" {
		HKEY = registry.CURRENT_USER
	} else if root == "HKEY_LOCAL_MACHINE" {
		HKEY = registry.LOCAL_MACHINE
	} else if root == "HKEY_USERS" {
		HKEY = registry.USERS
	} else if root == "HKEY_CURRENT_CONFIG" {
		HKEY = registry.CURRENT_CONFIG
	}

	return HKEY
}

func regKeyExists(keyroot, keypath string) int {

	var result int
	HKEY := getHKEY(keyroot)

	k, err := registry.OpenKey(HKEY, keypath, registry.QUERY_VALUE)

	if err != nil {
		//log.Fatal(err)
		result = 0
	} else {
		result = 1
	}

	defer k.Close()

	return result
}

func regWriteKey(keyroot, keypath string) {

	HKEY := getHKEY(keyroot)

	k, _, _ := registry.CreateKey(HKEY, keypath, registry.ALL_ACCESS)
	defer k.Close()
}

//设置注册表键值
func regWriteStringValue(keyroot, keypath, keyname, keyvalue string) int {

	var result int
	HKEY := getHKEY(keyroot)

	k, _, errhandle := registry.CreateKey(HKEY, keypath, registry.ALL_ACCESS)
	if errhandle != nil {
		return 2
	}
	defer k.Close()
	err := k.SetStringValue(keyname, keyvalue)
	if err != nil {
		result = 0
	} else {
		result = 1
	}

	return result
}

//删除注册表
func regDeleteKey(keyroot, keypath string) int {

	var result int
	HKEY := getHKEY(keyroot)

	err := registry.DeleteKey(HKEY, keypath)
	if err != nil {
		result = 0
	} else {
		result = 1
	}

	return result
}
