package operations

import (
	"HostSec/util"
	"bufio"
	"io"
	"os/exec"
)

func CcommandOpt(vectorcnname, command string) {
	util.PrintAttackResult(execCommand(command), vectorcnname)
}

func execCommand(command string) int {
	var res = 0

	cmd := exec.Command("cmd", "/c", command)
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
