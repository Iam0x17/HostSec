package main

import (
	"HostSec/cmd"
	"HostSec/control"
)

func main() {
	control.Load()
	cmd.Execute()
	control.Unload()
}
