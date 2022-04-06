package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sync"
	"unsafe"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func Gbk2Utf8Bytes(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func ExecCommand(cmd1 string, data chan []string, wg *sync.WaitGroup) {
	var allstr []string
	allstr = append(allstr, cmd1)
	cmd := exec.Command("cmd", "/c", cmd1)
	stderr, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		log.Println("exec the cmd ", cmd1, " failed")
	}
	cmdReader := io.MultiReader(stdout, stderr)
	scan := bufio.NewScanner(cmdReader)
	for scan.Scan() {
		s := scan.Text()
		ss, _ := Gbk2Utf8Bytes(Str2Bytes(s))
		//log.Print("build error: ", string(ss))
		//fmt.Println(string(ss))
		allstr = append(allstr, string(ss))
		//scan.WriteString(string(ss))
		//cmdReader.WriteString("\n")
	}
	//fmt.Println(allstr)
	//for _, v := range allstr {
	//	fmt.Println(v)
	//}
	allstr = append(allstr, "***********************************************")
	data <- allstr
	wg.Done()
}

func consume(data chan []string, done chan bool) {
	f, err := os.Create("collection.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		//fmt.Println(d)
		for _, v := range d {
			_, err = fmt.Fprintln(f, v)
			if err != nil {
				fmt.Println(err)
				f.Close()
				done <- false
				return
			}
		}

	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

func main() {
	var collection = []string{"systeminfo", "net user", "whoami"}
	wg := sync.WaitGroup{}
	data := make(chan []string)
	done := make(chan bool)

	for i := 0; i < len(collection); i++ {
		wg.Add(1)
		go ExecCommand(collection[i], data, &wg)
	}

	go consume(data, done)
	go func() {
		wg.Wait()
		close(data)
	}()
	d := <-done
	if d == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}

}
