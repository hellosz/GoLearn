// Package main provides ...
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 支持配置参数现实执行明细
var verbose = flag.Bool("v", false, "show verbose process message")

func main() {
	// dir := "/Users/patpat/User"
	// 初始化参数
	flag.Parse()

	// 初始化目录
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	// 遍历初始目录
	filesize := make(chan int64)

	// 每个目录使用一个goroutine，增加遍历速度
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go workdir(root, filesize, &wg)
	}

	// 关闭接收文件大小的通道
	go func() {
		wg.Wait()
		close(filesize)
	}()

	// 接收文件大小
	var nfilesize int64
	var nfile int64

	// 使用select多路复用，实现每500毫秒输出一次结果
loop:
	for {
		select {
		case size, ok := <-filesize:
			if !ok {
				break loop
			}

			nfilesize += size
			nfile++
		case _ = <-tick:
			printDiskUsage(nfilesize, nfile)
		}
	}

	// 输出结果
	printDiskUsage(nfilesize, nfile)
}

// printDiskUsage 输出结果明细
func printDiskUsage(filesize int64, file int64) {
	fmt.Printf("file number:%d, file size:%.1fMB\n", file, float64(filesize)/1e6)
}

// 使用token当作令牌桶进行限流，避免开启过多的文件
var token = make(chan struct{}, 20)

// dirents 返回目录信息
func dirents(dir string) []os.DirEntry {
	token <- struct{}{}
	entries, err := os.ReadDir(dir)
	_ = <-token
	if err != nil {
		log.Println(err)
		return nil
	}

	return entries
}

// workdir 遍历目录，统计文件大小
func workdir(dir string, filesize chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	entries := dirents(dir)
	for _, entry := range entries {
		if entry.IsDir() {
			// fmt.Printf("%s is dir\n", entry.Name())
			dirpath := filepath.Join(dir, entry.Name())

			// 递归查找子目录
			wg.Add(1)
			go workdir(dirpath, filesize, wg)
		} else {
			fileinfo, _ := entry.Info()
			// 返回文件大小
			filesize <- fileinfo.Size()
			// fmt.Printf("%s is file size(%d)\n", entry.Name(), fileinfo.Size())
		}

	}

}
