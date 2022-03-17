package database

type RegisterDB struct {
	VectorName   string `json:"vector_name"`
	VectorCnName string `json:"vector_cn_name"`
	KeyRoot      string `json:"key_root"`
	KeyPath      string `json:"key_path"`
	KeyName      string `json:"key_name"`
	KeyValue     string `json:"key_value"`
	OptType      string `json:"opt_type"`
}

type CommandDB struct {
	VectorName   string `json:"vector_name"`
	VectorCnName string `json:"vector_cn_name"`
	Command      string `json:"command"`
}

type FileDB struct {
	VectorName   string `json:"vector_name"`
	VectorCnName string `json:"vector_cn_name"`
	FilePath     string `json:"file_path"`
	FileContent  string `json:"file_content"`
	OptType      string `json:"opt_type"`
}

type VectorListDB struct {
	VectorName   string
	VectorCnName string
	Type         string
}

func (RegisterDB) TableName() string {
	return "registerdb"
}

func (FileDB) TableName() string {
	return "filedb"
}

func (CommandDB) TableName() string {
	return "commanddb"
}

func (VectorListDB) TableName() string {
	return "vectorlistdb"
}
