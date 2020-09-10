/*
浏览器自动打开网页
*/
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
	"path/filepath"
)

var (
	port int
	host *string
)

func init() {
	//两种方式绑定命令行参数
	flag.IntVar(&port, "p", 1212, "web port setting ")
	host = flag.String("host", "127.0.0.1", "web url setting")
	
	//覆盖默认 -help 描述
	flag.Usage = help

	//命令行参数解析
	flag.Parse()
}

//自定义帮助说明
func help() {
	fmt.Printf("[Usage] %s [-host] [-p] \n", filepath.Base(os.Args[0]))
	
	//默认参数说明打印
	flag.PrintDefaults()
}

//打开 url 是否成功判断,目前只支持 windows平台
//return bool
func open(url string) bool {
	args := []string{}
	switch runtime.GOOS {
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		return false
	}

	//构造命令
	command := exec.Command(args[0], append(args[1:], url)...)
	
	return command.Start() == nil && openSucc(command, 3)
}

//命令执行成功超时判断
func openSucc(cmd *exec.Cmd, t time.Duration) bool {
	//定义一个存储 error信息的管道
	errorChan := make(chan error, 1)
	go func() {
		errorChan <- cmd.Wait()
	}()

	select {
	case <-time.After(t)://超时处理
		return false
	case err := <-errorChan: //命令执行成功与否
		return err == nil
	}
}

func main() {
	//简单参数判断，不支持 0端口 及 空host
	if port == 0 || *host == "" {
		fmt.Println("Please input the right params")
		flag.Usage()
		
		os.Exit(1)
	}
	addr := fmt.Sprintf("%s:%d", *host, port)

	//另起一个 goroutine 用于浏览器自动打开url
	go func() {
		url := fmt.Sprintf("http://%s", addr)
		if open(url) {
			fmt.Printf("a browser is opening , if not ,please do it yourself and visti %s \n", url)
		} else {
			fmt.Printf("Please open your web browser and visit %s \n", url)
		}
	}()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Hello World!!!")
	})

	http.ListenAndServe(addr, nil)
}
