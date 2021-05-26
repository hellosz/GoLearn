package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	countingDown2()
	// workerSchedule()
}

// countingDown 发射倒计时
func countingDown() {
	fmt.Println("start counting down")
	tick := time.Tick(1 * time.Second)

	for counter := 10; counter > 0; counter-- {
		fmt.Printf("counting = %+v\n", counter)
		<-tick
	}

	launch()
}

// launch 火箭发射
func launch() {
	fmt.Println("launching...")
}

// countingDown2 可中断的倒计时发射器
func countingDown2() {
	fmt.Println("start counting down, press any key to stop it...")

	// 接收中止信号的通道
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	// 使用Ticker，避免gorouting泄漏
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	after := time.After(10 * time.Second)
	n := 10

Loop:
	for {
		select {
		case <-ticker.C:
			fmt.Printf("left %d second\n", n)
			if n == 1 {
				ticker.Stop()
			}
			n--
		case <-abort:
			fmt.Println("launch stopped")
			break Loop
		case <-after:
			fmt.Println("rocket launch success")
			break Loop
		}
	}
}

// workerSchedule 通过无缓存通道实现均匀调度
func workerSchedule() {
	fmt.Println("balanced worker schedule")

	ch := make(chan int, 1)
	for i := 10; i > 0; i-- {
		select {
		case val := <-ch:
			fmt.Printf("received val = %+v\n", val)
		case ch <- i:
			// TODO
			fmt.Printf("send val = %+v\n", i)

			// default:
			// TODO 实现非阻塞操作
		}
	}
}
