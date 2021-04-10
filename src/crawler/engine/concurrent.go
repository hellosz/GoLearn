package engine

import (
	// "fmt"
	"log"

	"hellosz.top/src/crawler/types"
)

// Concurrent 并发引擎
type Concurrent struct {
	Scheduler   Scheduler
	WorkerCount int
}

// Scheduler 调度器接口
type Scheduler interface {
	ReadyNotifier
	Submit(types.Request)
	WorkChan() chan types.Request
	Run()
}

// ReadyNotifier worker通知接口
type ReadyNotifier interface {
	WorkerReady(chan types.Request)
}

// Run 启动并发引擎
func (e *Concurrent) Run(seeds []types.Request) {
	// channel发送请求，channel接受结果
	out := make(chan types.ParseResult)
	e.Scheduler.Run()

	// 创建Worker
	for i := 0; i < e.WorkerCount; i++ {
		e.CreateWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}
	// 发送初始请求
	for _, seed := range seeds {
		e.Scheduler.Submit(seed)
	}

	// 处理返回值
	for {
		result := <-out

		for _, item := range result.Items {
			log.Printf("Got item %s\n", item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

// CreateWorker 创建工作者
func (Concurrent) CreateWorker(in chan types.Request, out chan types.ParseResult, notifier ReadyNotifier) {
	go func() {
		for {
			notifier.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
