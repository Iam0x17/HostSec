package models

type FileDB struct {
	VectorName   string `json:"vector_name"`
	VectorCnName string `json:"vector_cn_name"`
	FileFullPath string `json:"file_full_path"`
	FileContent  string `json:"file_content"`
	OptType      string `json:"opt_type"`
}

//定义结构体操作的数据库表
func (FileDB) TableName() string {
	return "filedb"
}
