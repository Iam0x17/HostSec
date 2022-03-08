package models

type RegisterDB struct {
	VectorName   string `json:"vector_name"`
	VectorCnName string `json:"vector_cn_name"`
	KeyRoot      string `json:"key_root"`
	KeyPath      string `json:"key_path"`
	KeyName      string `json:"key_name"`
	KeyValue     string `json:"key_value"`
	OptType      string `json:"opt_type"`
}

//定义结构体操作的数据库表
func (RegisterDB) TableName() string {
	return "registerdb"
}
