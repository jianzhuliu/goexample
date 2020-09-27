/*
调用os.Exit函数从任何函数里退出一个程序。
os.Exit函数调用接受一个int代码值做为参数并将此代码返回给操作系统
*/

package main

import "os"
import "time"

func main() {
	go func() {
		time.Sleep(time.Second)
		os.Exit(1)
	}()
	select {}
}
