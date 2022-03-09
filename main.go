package main

import (
	"HostSec/core"
	"HostSec/core/cmd"
)

func main() {
	core.Load()
	cmd.Execute()
	//vectorName := "MalDesktopLNK"
	//attackType := "file"
	//core.AttackSingle(vectorName, attackType)
	//core.AttackMulti()
	//AddData()
	core.Unload()
	//models.CreateDB()
	//models.AddData()

}
