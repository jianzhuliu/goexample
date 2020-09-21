/*
帮助文件
*/

package main

import (
	"os"
	"fmt"
	"runtime"
	"time"
	"os/exec"
)


//获取当前目录
func getwd() string{
	curPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return curPath
}

//创建不存在目录
func createDir(dir string){
	err := os.MkdirAll(dir, perm)
	if err != nil {
		fmt.Println("create dir fail,",dir,err)
	}
}

//自动打开url 
func openUrl(targetUrl string) bool {
	args := []string{}
	switch runtime.GOOS {
		case "windows":
			args = []string{"cmd", "/c", "start"}
		default:
			return false
	}
	
	cmd := exec.Command(args[0], append(args[1:], targetUrl)...)
	return cmd.Start() == nil && openUrlCheck(cmd, 3)
}

//url是否成功打开，带超时判断
func openUrlCheck(cmd *exec.Cmd, t time.Duration) bool {
	errChan := make(chan error, 1)
	go func(){
		errChan <- cmd.Wait()
	}()

	select {
		case <- time.After(t):
			return false 
		case err := <- errChan:
			return err == nil 
	}
}


