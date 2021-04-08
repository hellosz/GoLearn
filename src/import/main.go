package main

import (
	"fmt"
	"time"

	"hellosz.top/src/import/worker"
)

const workerCount = 20

func main() {

	// 扫描目录
	fmt.Println("读取目录")
	files := worker.ScanDir(worker.Dir)
	fileChan := make(chan worker.FileInfo, 5)

	// 创建工作者
	fmt.Println("创建goroutine")
	for i := 0; i < workerCount; i++ {
		worker.CreateWorker(fileChan)
	}

	fmt.Println("开始读取文件")
	for fileNum, file := range files {
		fmt.Printf("读取第%d个文件\n", fileNum+1)
		fileChan <- worker.FileInfo{
			FileName: file,
			FileNum:  fileNum + 1,
		}
	}

	// 沉睡一秒，保证数据全部读取完毕
	time.Sleep(1 * time.Second)
}
