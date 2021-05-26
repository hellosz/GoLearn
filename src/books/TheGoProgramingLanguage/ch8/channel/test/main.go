// Package main provides ...
package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("睡眠2秒结束, 关闭通道")
		close(ch)
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			ch <- "hello"
		}
	}()

	for {
		message, ok := <-ch
		fmt.Printf("%v, %t\n", message, ok)

		if ch == nil || ok == false {
			fmt.Println("通道关闭，退出循环")
			break
		}
	}
}
