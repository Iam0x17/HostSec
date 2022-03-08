package models

type CommandDB struct {
	VectorName   string `json:"vector_name"`
	VectorCnName string `json:"vector_cn_name"`
	Command      string `json:"command"`
}

//定义结构体操作的数据库表
func (CommandDB) TableName() string {
	return "commanddb"
}
