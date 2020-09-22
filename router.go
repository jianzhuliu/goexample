/*
路由定义
*/

package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"io/ioutil"
)

func init(){
	//定义路由
	//静态文件处理
	//http.Handle("/static/", http.FileServer(http.Dir(basePath)))
	http.Handle("/static/",http.StripPrefix("/static/", http.FileServer(http.Dir(staticPath))))
	http.Handle("/favicon.ico", http.FileServer(http.Dir(staticPath)))
	
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/demo/", handleDemo)
}

type IndexData struct {
	DemoItems []DemoItems
}

//处理首页
func handleIndex(rw http.ResponseWriter, r *http.Request){
	//列举出 demo 目录下所有一层文件夹及其下一层文件
	tmpData := IndexData{}
	tmpData.DemoItems = demoItems
	
	err := tmpl.ExecuteTemplate(rw, "index", tmpData)
	if err != nil {
		fmt.Println("handleIndex -- fail", err)
	}
}

type demoData struct{
	Content string
}

//处理单个文件
func handleDemo(rw http.ResponseWriter, r *http.Request){
	urlPath := r.URL.Path
	//构造url
	file := filepath.Join(basePath,urlPath)
	
	//读取内容
	content, err := ioutil.ReadFile(file)
	
	if err != nil {
		fmt.Println("handleDemo -- fail --",urlPath, err)
		http.NotFound(rw, r)
		return 
	}
	
	tmpData := demoData{}
	tmpData.Content = string(content)
	
	err = tmpl.ExecuteTemplate(rw, "content", tmpData)
	if err != nil {
		fmt.Println("handleDemo -- fail", err)
	}
	
}

