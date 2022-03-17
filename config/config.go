package config

var (
	AttackDB = []string{"register", "file", "command"}
)

//DBPath
const (
	HipsJsonDir = `json\hipsjson`
	HipsDBName  = `hips.db`
)

//AttackType
const (
	HipsReg  = "register"
	HipsFile = "file"
	HipsCmd  = "command"
)
