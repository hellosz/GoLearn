// Package main provides ...
package main

import (
	"fmt"
	"time"
)

func main() {
	testSelect()
}

func testSelect() {
	ccClose := make(chan int, 2)
	go func() {
		for i := 0; i < 2; i++ {
			ccClose <- i
		}
		time.Sleep(time.Second)
		close(ccClose)

	}()

	ssClose := make(chan string, 2)
	go func() {
		tempString := []string{"hello", "world"}
		for _, str := range tempString {
			ssClose <- str
		}
		time.Sleep(time.Second)
		close(ssClose)
	}()

	ccClosed := false
	ssClosed := false

	for {
		// 如果全部关闭，则退出
		if ccClosed && ssClosed {
			fmt.Println("通道关闭，退出循环...")
			return
		}

		select {
		case num, ok := <-ccClose:
			if !ok {
				ccClosed = true
				break // 提前终止select
			}
			fmt.Printf("ccClose channel received:%d\n", num)
		case str, ok := <-ssClose:
			if !ok {
				ssClosed = true
				break // 提前终止select
			}
			fmt.Printf("ssClose channel received:%s\n", str)
		default:
			fmt.Println("waiting for....")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
