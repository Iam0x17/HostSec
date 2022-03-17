package main

import (
	"HostSec/cmd"
	"HostSec/controller"
)

func main() {
	controller.Load()
	cmd.Execute()
	controller.Unload()
}
