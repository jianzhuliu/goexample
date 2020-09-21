/*
入口文件，支持配置端口及IP
demo:	go run . -port=1232 -host=127.0.0.1
*/

package main

import (
	"fmt"
	"flag"
	"os"
	
	"path/filepath"
	"net/http"
	"html/template"
)

var (
	port int 
	host *string
	
	//定义路径
	basePath string
	demoPath string
	tmplPath string
	staticPath string
	logsPath string
	
	//默认文件及目录创建权限
	perm os.FileMode = 0644
	
	tmpl *template.Template
	
)

//初始化
func init(){
	flag.IntVar(&port,"port",1314,"set port")
	host = flag.String("host","127.0.0.1","set host") 
	
	flag.Usage = usage //自定义说明行数
	
	initPath() //初始化路径
	
	initTmplFiles() //加载模板文件
	
	LoadDemoItems() //加载 demo 文件夹列表
}

//自定义说明行数
func usage(){
	fmt.Printf("Usage of %s:\n", filepath.Base(os.Args[0]))
	flag.PrintDefaults()
}

//初始化路径
func initPath(){
	basePath = getwd()
	demoPath = filepath.Join(basePath,"demo")
	tmplPath = filepath.Join(basePath,"template")
	staticPath = filepath.Join(basePath,"static")
	logsPath = filepath.Join(basePath,"logs")
	
	//logs 目录不存在，则自动创建
	createDir(logsPath)
}

//加载模板文件 
func initTmplFiles(){
	pattern := filepath.Join(tmplPath, string(os.PathSeparator), "*.tmpl")
	tmpl = template.Must(template.ParseGlob(pattern))
}

func main(){
	flag.Parse()
	
	//简单参数验证
	if port <= 0 || len(*host) == 0 {
		fmt.Println("please input the right params")
		os.Exit(1)
	}
	
	addr := fmt.Sprintf("%s:%d",*host, port)
	
	go func(){
		targetUrl := fmt.Sprintf("http://%s", addr)
		
		if openUrl(targetUrl) {
			fmt.Printf("a browser is open, if not , please visit %s \n", targetUrl)
		} else {
			fmt.Printf("Please open your web browser and visit %s \n",targetUrl)
		}
	}()
	
	if err := http.ListenAndServe(addr,nil); err != nil {
		fmt.Println("start failed -- ", err)
	}
}



