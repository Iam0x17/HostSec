package models

type RegisterDB struct {
	VectorName   string
	VectorCnName string
	KeyRoot      string
	KeyPath      string
	KeyName      string
	KeyValue     string
	OptType      string
}

//定义结构体操作的数据库表
func (RegisterDB) TableName() string {
	return "registerdb"
}
