package main

import (
	"HostSec/core"
	"HostSec/core/cmd"
	"HostSec/util"
)

func main() {
	cmd.Execute()
	util.IsElevated()
	//vectorName := "MalDesktopLNK"
	//attackType := "file"
	//core.AttackSingle(vectorName, attackType)
	//core.AttackMulti()
	//AddData()
	core.Unload()
	//models.CreateDB()
	//models.AddData()

}
