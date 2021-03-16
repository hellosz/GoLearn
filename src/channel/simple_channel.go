package main

import (
	"fmt"
	"time"
)

func main() {
	ch := doWork()
	tick := time.Tick(time.Second * 20)

	for {
		select {
		case sign := <-ch:
			fmt.Println("Received channel data is", sign)
		case <-tick:
			fmt.Println("Time ended, stop the process")
			return
		default:
			fmt.Println("No Received data")
		}

		time.Sleep(500 * time.Millisecond)
	}
}

// 创建工作者，并且返回
func doWork() chan int {
	ch := make(chan int)

	go func() {
		i := 0
		for {
			ch <- i
			i++

			// 休眠0.5秒
			time.Sleep(time.Millisecond * 1000)
		}

	}()
	return ch
}
