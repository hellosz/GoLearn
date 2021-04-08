package worker

import (
	"fmt"
	"io/ioutil"
	"log"

	"hellosz.top/src/import/connetion"
	"hellosz.top/src/import/parser"
	"hellosz.top/src/import/utils"

	// "hellosz.top/src/import/connetion"
	_ "hellosz.top/src/import/model"
)

// 文件信息
type FileInfo struct {
	FileName string
	FileNum  int
}

const Dir = "/Users/patpat/Documents/Dingding/mms-s3-access-logs/mms"

// ScanDir 读取目录下的文件
func ScanDir(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("%s\n", err)

	}

	var fileNames []string
	for _, file := range files {
		// fmt.Println(file.Name())
		fileNames = append(fileNames, file.Name())
	}

	fmt.Println(len(files))

	return fileNames
}

// ReadFile 读取文件
func ReadFile(file string) []byte {
	filepath := Dir + "/" + file

	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf("读取文件:%s失败,失败原因:%v", filepath, err)
	}

	return content
}

// CreateWorker 建工作者，去读取文件
func CreateWorker(fileinfo chan FileInfo, config utils.Config) {
	// 获取数据库连接
	client := connetion.Get(config.Connection)

	go func() {
		for {
			file := <-fileinfo
			// 读取目录下文件J
			content := ReadFile(file.FileName)

			// 解析日志到Model中j
			access := parser.Log(content)

			// 保存model中的数据
			id, err := access.Create(client)
			if err != nil {
				log.Printf("保存第%d条数据%v失败, 失败原因:%v", file.FileNum, access, err)
			} else {
				log.Printf("保存第%d条数据:%s成功, 返回结果%d", file.FileNum, access.BucketKey, id)
			}
		}
	}()
}
