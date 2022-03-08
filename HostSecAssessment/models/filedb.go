package models

type FileDB struct {
	VectorName   string
	VectorCnName string
	FileFullPath string
	FileContent  string
	OptType      string
}

//定义结构体操作的数据库表
func (FileDB) TableName() string {
	return "filedb"
}
