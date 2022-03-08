package models

type CommandDB struct {
	VectorName   string
	VectorCnName string
	Command      string
}

//定义结构体操作的数据库表
func (CommandDB) TableName() string {
	return "commanddb"
}
