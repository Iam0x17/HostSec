package hips

import (
	"bufio"
	"io"
	"os/exec"
)

type CommandSrv interface {
	ExecCommand() int
}

type CommandData struct {
	Command string
}

func NewCommandVector(command string) CommandSrv {
	commandData := CommandData{
		Command: command,
	}

	return commandData
}

func (command CommandData) ExecCommand() int {
	var res = 0

	cmd := exec.Command("cmd", "/c", command.Command)
	stderr, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		//log.Println("exec the cmd ", command, " failed")
		res = 2
	}
	cmdReader := io.MultiReader(stdout, stderr)
	scan := bufio.NewScanner(cmdReader)
	for scan.Scan() {
		//s := scan.Text()
		//ss, _ := util.Gbk2Utf8Bytes(util.Str2Bytes(s))
		//log.Print("build error: ", string(ss))
		//errBuf.WriteString(string(ss))
		//errBuf.WriteString("\n")
		res = 1
	}
	return res
}
