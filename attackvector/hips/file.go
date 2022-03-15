package hips

import (
	"HostSec/util"
	"os"
)

type FileSrv interface {
	WriteFile() int
	DelFile() int
}

type FileData struct {
	FilePath    string
	FileContent string
	OptType     string
}

func NewFileVector(filepath, filecontent, opttype string) FileSrv {
	fileData := FileData{
		FilePath:    filepath,
		FileContent: filecontent,
		OptType:     opttype,
	}

	return fileData
}

func (file FileData) WriteFile() int {
	var flag int
	switch file.OptType {
	case "create":
		flag = os.O_WRONLY | os.O_CREATE
		break
	case "write_append":
		flag = os.O_WRONLY | os.O_APPEND
		break
	}
	realPath := util.GetRealPath(file.FilePath)

	hFile, err := os.OpenFile(realPath, flag, 0666)
	if err != nil {
		//fmt.Printf("open file err=%v\n", err)
		return 0
	}
	defer hFile.Close()

	_, errWrite := hFile.WriteString("\r\n" + file.FileContent)
	if errWrite != nil {
		//fmt.Printf("write file err=%v\n", err)
		return 0
	}

	return 1

}

func (file FileData) DelFile() int {
	checkFileIsExist(file.FilePath)
	err := os.Remove(file.FilePath)
	if err != nil {
		return 0
	} else {
		return 1
	}
}

//判断文件是否存在
func checkFileIsExist(filepath string) bool {
	var exist = true
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
