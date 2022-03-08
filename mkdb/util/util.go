package util

import (
	"io/ioutil"
	"strings"
)

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
