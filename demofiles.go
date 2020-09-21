/*
文件列表
*/

package main 

import(
	"fmt"
	"os"
	"sort"
	"sync"
	"path/filepath"
)

type DemoItem struct{
	Name string
}

type DemoItems struct{
	Name string
	Items []DemoItem
}

var demoItems []DemoItems
var once sync.Once

func LoadDemoItems(){
	once.Do(func(){
		tmpData, err := loadDemoItems()
		if err != nil {
			fmt.Println("loadDemoItems -- fail -- ", err)
		}
		
		demoItems = tmpData
	})
}

func loadDemoItems() ([]DemoItems,error){

	f, err := os.Open(demoPath)
	if err != nil {
		return nil , err 
	}
	
	//读取文件夹列表
	names, err := f.Readdirnames(-1)
	f.Close()
	
	if err != nil {
		return nil, err
	}
	
	sort.Strings(names)
	var result []DemoItems
	
	for _, name := range names {
		tmpdir := filepath.Join(demoPath, name)
		file, err := os.Open(tmpdir)
		if err != nil {
			continue
		}
		
		//读取文件夹下文件列表
		fileInfos, err := file.Readdir(-1)
		file.Close()
		
		if err != nil {
			continue
		}
			
		tmpItems := DemoItems{}
		flagHasFile := false 
		
		for _, fileinfo := range fileInfos {
			if !fileinfo.IsDir() {
				//读取文件
				tmpItem := DemoItem{}
				tmpItem.Name = fileinfo.Name()
				
				tmpItems.Items = append(tmpItems.Items, tmpItem)
				flagHasFile = true 
			}
		}

		if flagHasFile {
			tmpItems.Name = name 
			result = append(result, tmpItems)
		}
	}
	
	return result, nil 
}