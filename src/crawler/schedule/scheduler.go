package schedule

import "hellosz.top/src/crawler/types"

// Scheduler 并发调度器
type Scheduler struct {
	chanIn chan types.Request
}

// Submit 向通道发送请求
func (s *Scheduler) Submit(request types.Request) {
	go func() {
		s.chanIn <- request
	}()
}

// ConfigureMasterChanIn 配置输入通道
func (s *Scheduler) ConfigureMasterChanIn(chanIn chan types.Request) {
	s.chanIn = chanIn
}
