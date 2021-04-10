package schedule

import "hellosz.top/src/crawler/types"

// Simple 并发调度器
type Simple struct {
	chanIn chan types.Request
}

// Submit 向通道发送请求
func (s *Simple) Submit(request types.Request) {
	go func() {
		s.chanIn <- request
	}()
}

// ConfigureMasterChanIn 配置输入通道
func (s *Simple) ConfigureMasterChanIn(chanIn chan types.Request) {
	s.chanIn = chanIn
}
