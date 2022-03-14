package operations

import (
	"HostSec/util"
	"os"
)

func FileOpt(vectorcnname, filefullpath, filecontent, opttype string) {
	var res int

	switch opttype {
	case "create", "write_append":
		res = writeFile(filefullpath, filecontent, opttype)
		break
	case "del":
		res = delFile(filefullpath)
		break
	}
	util.PrintAttackResult(res, vectorcnname)
}

//判断文件是否存在
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func writeFile(filefullpath, filecontent, opttype string) int {
	var flag int
	switch opttype {
	case "create":
		flag = os.O_WRONLY | os.O_CREATE
		break
	case "write_append":
		flag = os.O_WRONLY | os.O_APPEND
		break
	}
	realPath := util.GetRealPath(filefullpath)

	file, err := os.OpenFile(realPath, flag, 0666)
	if err != nil {
		//fmt.Printf("open file err=%v\n", err)
		return 0
	}
	defer file.Close()

	_, errWrite := file.WriteString("\r\n" + filecontent)
	if errWrite != nil {
		//fmt.Printf("write file err=%v\n", err)
		return 0
	}

	return 1

}

func delFile(filefullpath string) int {
	checkFileIsExist(filefullpath)
	err := os.Remove(filefullpath)
	if err != nil {
		return 0
	} else {
		return 1
	}
}
