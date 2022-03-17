package winpath

import "strings"

func GetWinPath(path string) string {
	sep := "%"
	if strings.LastIndex(path, sep) == -1 {
		return path
	}
	var envPath string
	pathArray := strings.Split(path, sep)
	switch pathArray[1] {
	case "Desktop":
		envPath, _ = Desktop()
		break
	case "Favorites":
		envPath, _ = Favorites()
	case "AppData":
		envPath, _ = CommonAppData()
		break
	default:
		return path
	}
	realPath := envPath + pathArray[2]
	return realPath
}
