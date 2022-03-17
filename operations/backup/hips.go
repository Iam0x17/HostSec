package backup

import (
	"HostSec/attackvector/hips"
	"strings"
)

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
	dstFilePath := `D:\HostSec\backup\` + str[len(str)-1]
	return dstFilePath
}
