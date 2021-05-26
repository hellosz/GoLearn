package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
)

func main() {
	// 构建文件列表
	times := 20
	filenames := BuildFileSlice(times)
	// fmt.Println(filenames)

	// 压缩文件
	files, err := MakeThumnails5(filenames)

	// 返回结果异常处理
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(files)
}

// 	ImageFile 压缩文件
func ImageFile(name string) (string, error) {
	if strings.Contains(name, "20.png") {
		return "", errors.New("文件包含20.png，出现异常")
	}

	fmt.Printf("正在压缩%s文件...\n", name)
	fmt.Printf("文件%s压缩成功...\n", name)
	return "zip" + name, nil
}

// 随机返回文件大小
func ImageSize(name string) int64 {
	return rand.Int63()
}

// BuildFileSlice 构建文件切片
func BuildFileSlice(times int) []string {
	filenameSuffix := ".png"
	var filenames []string
	for i := 1; i <= times; i++ {
		filenames = append(filenames, strconv.Itoa(i)+filenameSuffix)
	}
	return filenames
}

// MakeThumnails 普通的压缩文件的方法
func MakeThumnails(filenames []string) error {
	for _, f := range filenames {
		if _, err := ImageFile(f); err != nil {
			log.Printf("文件:%s 压缩失败，原因:%s", f, err.Error())
		}
	}

	return nil
}

// 使用go关键字并发(会出现文件没有压缩完，但是主程序退出的情况)
func MakeThumnails2(filenames []string) error {
	for _, f := range filenames {
		go ImageFile(f)
	}

	return nil
}

// MakeThumnails3 使用通道通信，控制程序结束
func MakeThumnails3(filenames []string) error {
	ch := make(chan struct{})

	// 压缩文件
	for _, f := range filenames {
		go func(f string) {
			ImageFile(f)
			ch <- struct{}{}
		}(f)
	}

	// 等待文件执行结束
	for range filenames {
		<-ch
	}

	return nil
}

// MakeThumnails4 允许返回错误，中断程序执行
func MakeThumnails4(filenames []string) error {
	errChan := make(chan error, len(filenames))

	// 压缩文件
	for _, f := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			errChan <- err
		}(f)
	}

	// 等待文件执行结束
	for range filenames {
		if err := <-errChan; err != nil {
			return err
		}
	}

	return nil
}

// MakeThumnails5 以默认顺序返回压缩文件名，如果出现错误则中断程序返回报错
func MakeThumnails5(filenames []string) ([]string, error) {
	type item struct {
		ImageFile string
		Err       error
	}

	// 返回结果通道
	ch := make(chan item, len(filenames))

	// 压缩文件
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.ImageFile, it.Err = ImageFile(f)
			ch <- it
		}(f)
	}

	// 等待文件执行结果，有报错直接返回
	imageFiles := []string{}
	for range filenames {
		it := <-ch
		if it.Err != nil {
			return nil, it.Err
		}

		fmt.Println("returned filename :" + it.ImageFile)

		imageFiles = append(imageFiles, it.ImageFile)
	}

	return imageFiles, nil
}

// MakeThumnails6 通过通道传递文件名，返回所有压缩后的文件大小
func MakeThumnails6(filenames chan string) int64 {
	// 收集压缩文件大小
	size := make(chan int64)
	wg := sync.WaitGroup{}

	// 压缩文件
	for filename := range filenames {
		wg.Add(1)
		compressFile, err := ImageFile(filename)
		if err != nil {
			log.Fatal(err)
			return 0
		}

		size <- ImageSize(compressFile)
		wg.Done()
	}

	// 使用goroutine关闭size通过
	go func() {
		wg.Wait()
		close(size)
	}()

	// 等待文件执行，收集压缩文件大小
	var totalSize int64
	for singleSize := range size {
		totalSize += singleSize
	}

	// 返回结果
	return totalSize
}
