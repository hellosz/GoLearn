// Package  provides ...
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	schedule()
}

// schedule 并发爬虫任务调度方法
func schedule() {
	// 使用channel实现队列
	worklist := make(chan []string)

	// 发送到worklist中的消息数量
	var n int = 0
	n++

	// 使用命令行参数初始化channel
	go func() {
		arguments := os.Args[1:]

		fmt.Printf("initial arguments is:%v\n", arguments)
		worklist <- arguments
	}()

	// 广度优先遍历
	seenLinks := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			// 爬取链接
			if !seenLinks[link] {
				seenLinks[link] = true

				// 消息数量加1
				n++

				// 爬取链接
				go func(link string) {
					list, _ := crawl(link)
					fmt.Printf("crawler link: %s result is :%v\n", link, list)

					worklist <- list
				}(link)
			}
		}
	}

}

var token = make(chan struct{}, 20)

// crawl 爬取指定链接，并返回结果；通过计数信号量增加限流限制
func crawl(link string) ([]string, error) {
	// 打印链接
	fmt.Printf("正在爬取链接link = %+v\n", link)

	// 拿令牌
	token <- struct{}{}
	// 爬取链接
	list, err := extract(link)

	// 释放令牌
	<-token

	if err != nil {
		// 异常处理，输出日志，并返回错误
		log.Fatal(err)
		return []string{}, err
	}

	return list, nil
}

// extract 链接爬取核心方法
func extract(link string) ([]string, error) {
	// 休眠随机时间
	// randomInt := rand.Intn(3)
	randomInt := 1
	time.Sleep(time.Duration(randomInt) * time.Second)

	return []string{link}, nil
}

// 自定义的range方法
func customRange() {
	// 遍历一个数组/array
	arr := []string{"1.png", "2.png", "3.png", "4.png"}
	// 遍历时有两个返回值，第一个为key值，第二个为value值
	for idx, val := range arr {
		fmt.Printf("key:%d, value:%s", idx, val)
	}

	// 遍历一个channel
	ch := make(chan []string)
	go func() {
		fmt.Println("向channel发送消息")
		ch <- arr
		close(ch)
	}()

	// 遍历去接收channel中的值，直到关闭
	for arr := range ch {
		fmt.Println("从channel接收消息")
		fmt.Println(arr)
	}
}
