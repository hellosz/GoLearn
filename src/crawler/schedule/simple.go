package schedule

import "hellosz.top/src/crawler/types"

// Simple 并发调度器
type Simple struct {
	ChanIn chan types.Request
}

// Submit 向通道发送请求
func (s *Simple) Submit(request types.Request) {
	go func() {
		s.ChanIn <- request
	}()
}

// WorkChan 每个worker返回一个chan
func (s *Simple) WorkChan() chan types.Request {
	return s.ChanIn
}

// WorkerReady 将准备好的worker推送到队列
func (s *Simple) WorkerReady(worker chan types.Request) {
}

// Run 任务调度中心
func (s *Simple) Run() {
	s.ChanIn = make(chan types.Request)
}
