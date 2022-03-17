package hips

import (
	"HostSec/util"
	"fmt"
	"io"
	"os"
)

type FileSrv interface {
	WriteFile() int
	DelFile(dstfilepath ...string) int
	CopyFile(dstfilepath string) int
	DelCopy(dstfilepath string) int
}

type FileData struct {
	FilePath    string
	FileContent string
	OptType     string
}

func NewFileVector(filepath, filecontent, opttype string) FileSrv {
	fileData := FileData{
		FilePath:    util.GetRealPath(filepath),
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
	//realPath := util.GetRealPath(file.FilePath)

	hFile, err := os.OpenFile(file.FilePath, flag, 0666)
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

func (file FileData) DelFile(dstfilepath ...string) int {
	var err error
	if len(dstfilepath) == 0 {
		err = os.Remove(file.FilePath)
	} else {
		for _, v := range dstfilepath {
			realPath := util.GetRealPath(v)
			err = os.Remove(realPath)
		}
	}

	if err != nil {
		return 0
	} else {
		return 1
	}
}

func (file FileData) CopyFile(dstfilepath string) int {

	srcFile, errSrcFile := os.Open(file.FilePath)
	if errSrcFile != nil {
		fmt.Printf("open file err = %v\n", errSrcFile)
		return 0
	}

	defer srcFile.Close()
	//打开dstFileName
	dstFile, errDstFile := os.OpenFile(dstfilepath, os.O_WRONLY|os.O_CREATE, 0755)
	if errDstFile != nil {
		fmt.Printf("open file err = %v\n", errDstFile)
		return 0
	}
	defer dstFile.Close()
	_, errCopy := io.Copy(dstFile, srcFile)
	if errCopy != nil {
		fmt.Printf("open file err = %v\n", errCopy)
		return 0
	}

	return 1
}

func (file FileData) DelCopy(dstfilepath string) int {
	realPath := util.GetRealPath(dstfilepath)
	err := os.Remove(realPath)
	if err != nil {
		return 0
	}
	srcFile, errSrcFile := os.Open(file.FilePath)
	if errSrcFile != nil {
		fmt.Printf("open file err = %v\n", errSrcFile)
		return 0
	}

	defer srcFile.Close()
	//打开dstFileName
	dstFile, errDstFile := os.OpenFile(dstfilepath, os.O_WRONLY|os.O_CREATE, 0755)
	if errDstFile != nil {
		fmt.Printf("open file err = %v\n", errDstFile)
		return 0
	}
	defer dstFile.Close()
	_, errCopy := io.Copy(dstFile, srcFile)
	if errCopy != nil {
		fmt.Printf("open file err = %v\n", errCopy)
		return 0
	}

	return 1
}

func delFile(filepath string) int {
	//checkFileIsExist(file)
	err := os.Remove(filepath)
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
