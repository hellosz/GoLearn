package main

import (
	"flag"
	"fmt"
	"hellosz.top/src/import/utils"
	"hellosz.top/src/import/worker"
	"sync"
)

const workerCount = 20

func main() {
	// 读取配置文件
	configPath := flag.String("path", "config.yaml", "yml配置文件的地址")
	flag.Parse()

	config := utils.Config{}
	realconfig := config.Get(*configPath)

	// 扫描目录
	fmt.Println("读取目录")
	files := worker.ScanDir(realconfig.LogDir)
	fileChan := make(chan worker.FileInfo, 5)

	// 创建waitgroup
	wg := sync.WaitGroup{}
	wg.Add(len(files))

	// 创建工作者
	fmt.Println("创建goroutine")
	for i := 0; i < realconfig.WorkerCount; i++ {
		worker.CreateWorker(fileChan, *realconfig, &wg)
	}

	fmt.Println("开始读取文件")
	for fileNum, file := range files {
		fmt.Printf("读取第%d个文件\n", fileNum+1)
		fileChan <- worker.FileInfo{
			FileName: file,
			FileNum:  fileNum + 1,
		}
	}

	// 等待数据保存完成
	wg.Wait()
}
