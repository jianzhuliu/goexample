/*
Playground 相关
*/

package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"encoding/json"
	"go/format"
	"io/ioutil"
	"net/http"
	"os/exec"
	"path/filepath"
)

func init() {
	//初始化定义路由
	http.HandleFunc("/playground", handlePlayground)

	http.HandleFunc("/run", handleRun)
	http.HandleFunc("/fmt", handleFmt)
	http.HandleFunc("/imports", handleImports)
}

//demo 定义
var playgroundDemo string = `package main
import "fmt"
func main(){
	fmt.Println("Hello World !!!")
}
`

var errEmpty = errors.New("post data is empty")

//playground 主界面
func handlePlayground(rw http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(rw, "playground", playgroundDemo)

	if err != nil {
		fmt.Println("handlePlayground -- fail", err)
	}
}

//可以选择的控制字段有三种：
// `json:"-"`  			-：不要解析这个字段
// `json:",omitempty"`	omitempty：当字段为空（默认值）时，不要解析这个字段。
// `json:"sname"`  		FieldName：当解析 json 的时候，使用这个名字
type PlaygroundData struct {
	Error string `json:",omitempty"`
	Body  string `json:",omitempty"`
}

//json 返回数据
func JSON(rw http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")

	rw.WriteHeader(statusCode)

	encoder := json.NewEncoder(rw)
	if err := encoder.Encode(data); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

}

//错误信息统一处理
func handleError(rw http.ResponseWriter, r *http.Request, err error) {
	tmpData := PlaygroundData{}
	tmpData.Error = err.Error()
	JSON(rw, r, http.StatusOK, tmpData)
}

//统一正确数据回包
func handleOK(rw http.ResponseWriter, r *http.Request, body string) {
	tmpData := PlaygroundData{}
	tmpData.Body = body
	JSON(rw, r, http.StatusOK, tmpData)
}

//运行
func handleRun(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handleError(rw, r, errors.New("http method is not right"))
		return
	}

	//传递内容不能为空
	body := r.PostFormValue("body")
	if len(body) == 0 {
		handleError(rw, r, errEmpty)
		return
	}

	//创建临时目录
	tempDir, err := ioutil.TempDir("", "run")
	if err != nil {
		handleError(rw, r, err)
		return
	}

	//延迟删除目录及其下所有
	defer os.RemoveAll(tempDir)

	//目标文件
	file := filepath.Join(tempDir, "prog.go")

	//写入客户端上报内容
	err = ioutil.WriteFile(file, []byte(body), perm)
	if err != nil {
		handleError(rw, r, err)
		return
	}

	//带超时的上下文信息
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//先执行 go vet 进行语法之类的检查
	cmd := exec.CommandContext(ctx, "go", "vet", file)
	output, err := cmd.Output()
	if err != nil {
		handleError(rw, r, err)
		return
	}

	//说明检查出异常
	if len(output) > 0 {
		handleError(rw, r, errors.New(string(output)))
		return
	}

	//然后执行 go run
	output, err = exec.CommandContext(ctx, "go", "run", file).CombinedOutput()
	if err != nil {
		handleError(rw, r, err)
		return
	}

	handleOK(rw, r, string(output))
}

//fmt 格式化
func handleFmt(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handleError(rw, r, errors.New("http method is not right"))
		return
	}

	//简单参数判断
	body := r.PostFormValue("body")
	if len(body) == 0 {
		handleError(rw, r, errEmpty)
		return
	}

	content, err := format.Source([]byte(body))
	if err != nil {
		handleError(rw, r, err)
		return
	}
	/*
		//以下内容废弃
		//创建临时目录
		tempDir,err := ioutil.TempDir("", "fmt")
		if err != nil {
			handleError(rw,r, err)
			return
		}

		//延迟删除目录及其下所有
		defer os.RemoveAll(tempDir)

		//目标文件
		file := filepath.Join(tempDir, "prog.go")

		//写入客户端上报内容
		err = ioutil.WriteFile(file,[]byte(body), perm)
		if err != nil {
			handleError(rw,r, err)
			return
		}

		//上下文信息
		ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
		defer cancel()

		//执行格式化命令
		err = exec.CommandContext(ctx,"go","fmt",file).Run()
		if err != nil {
			handleError(rw,r, err)
			return
		}

		//成功后，读取格式化后的内容
		content, err := ioutil.ReadFile(file)
		if err != nil {
			handleError(rw,r, err)
			return
		}
		//*/

	handleOK(rw, r, string(content))
}

//自动导入
func handleImports(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handleError(rw, r, errors.New("http method is not right"))
		return
	}

	//简单参数判断
	body := r.PostFormValue("body")
	if len(body) == 0 {
		handleError(rw, r, errEmpty)
		return
	}

	handleError(rw, r, errors.New("not done yet"))
}
