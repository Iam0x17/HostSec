package util

/*
#include <permission.h>

*/
import "C"
import (
	"bytes"
	"fmt"
	"github.com/theherk/winpath"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
	"unsafe"
)

func IsElevated() {
	isAdmin := C.IsElevated()
	if isAdmin != 0 {
		fmt.Println("管理员权限")
	} else {
		fmt.Println("非管理员权限，请使用管理员权限，否则可能造成某些测试项不生效")
	}
}

func PrintAttackResult(res int, vectorcnname string) {
	var describe string
	switch res {
	case 0:
		describe = "[Success]:" + vectorcnname + " 攻击失败，主机防御成功"
		gResult.Success++
		break
	case 1:
		describe = "[Failed]:" + vectorcnname + " 攻击成功，主机防御失败"
		gResult.Failed++
		break
	case 2:
		describe = "[Denied]:" + vectorcnname + " 权限不足"
		gResult.Denied++
		break
	}
	ColorPrint(res, describe)
	writeLog(describe)
}

func GetRealPath(path string) string {
	sep := "%"
	if strings.LastIndex(path, sep) == -1 {
		return path
	}
	var envPath string
	pathArray := strings.Split(path, sep)
	switch pathArray[1] {
	case "Desktop":
		envPath, _ = winpath.Desktop()
		break
	default:
		return path
	}
	realPath := envPath + pathArray[2]
	return realPath
}

func Gbk2Utf8Bytes(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func GetFileList(folder string) []string {
	var fileList []string
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		if file.IsDir() {
			GetFileList(folder + "/" + file.Name())
		} else {
			//fmt.Println(folder + "/" + file.Name())
			file := folder + "/" + file.Name()
			fileList = append(fileList, file)
		}
	}
	return fileList
}

func ExistWhoStr(str string, strsilce []string) (bool, string) {
	for _, v := range strsilce {
		if strings.Contains(str, v) == true {
			return true, v
		}
	}
	return false, ""
}
