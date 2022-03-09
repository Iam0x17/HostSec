package util

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Result struct {
	Success int
	Failed  int
	Denied  int
}

var gResult = new(Result)
var gLogFileHandle *os.File
var signLog = false

func SetLogSign(sign bool) {
	if sign {
		fmt.Println("日志开启")
		createLogFile()
	}
	signLog = sign
}

func GetLogSign() bool {
	return signLog
}

func createLogFile() {
	var err error
	gLogFileHandle, err = os.OpenFile(getLogName(), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("创建日志文件失败")
	}
}

func getLogName() string {
	timeStr := time.Now().Format("2006-01-02_15-04-05") //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	return timeStr + ".log"
}

func writeLog(content string) {
	_, errWrite := gLogFileHandle.WriteString(content + "\r\n")
	if errWrite != nil {
		fmt.Println("日志写入失败")
	}
}

func WriteLogResult() {
	var conent string
	conent = "\r\n结果统计\r\n" + "防御成功:" + strconv.Itoa(gResult.Success) + " 防御失败:" + strconv.Itoa(gResult.Failed) + " 权限不足:" + strconv.Itoa(gResult.Denied)
	//conent = "结果\r\n" + "防御成功:" + string(gResult.Success) + "防御失败:" + string(gResult.Failed) + "权限不足:" + string(gResult.Denied)
	writeLog(conent)
	gLogFileHandle.Close()
}
