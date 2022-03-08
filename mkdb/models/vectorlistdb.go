package models

type VectorListDB struct {
	VectorName   string
	VectorCnName string
	Type         string
}

//定义结构体操作的数据库表
func (VectorListDB) TableName() string {
	return "vectorlistdb"
}
