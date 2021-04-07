package engine

import (
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
	Submit(types.Request)
	ConfigureMasterChanIn(chan types.Request)
}

// Run 启动并发引擎
func (e *Concurrent) Run(seeds []types.Request) {
	// channel发送请求，channel接受结果
	in := make(chan types.Request)
	out := make(chan types.ParseResult)
	e.Scheduler.ConfigureMasterChanIn(in)

	// 创建Worker
	for i := 0; i < e.WorkerCount; i++ {
		e.CreateWorker(in, out)
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
func (Concurrent) CreateWorker(in chan types.Request, out chan types.ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
