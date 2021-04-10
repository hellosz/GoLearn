package schedule

import "hellosz.top/src/crawler/types"

// Queue 并发调度器
type Queue struct {
	RequestChan chan types.Request
	WorkerChan  chan chan types.Request
}

// Submit 将准备的请求推送给队列
func (s *Queue) Submit(request types.Request) {
	go func() {
		s.RequestChan <- request
	}()
}

// ConfigureMasterChanIn 配置输入通道
func (s *Queue) ConfigureMasterChanIn(chanIn chan types.Request) {
	panic("deprecated method")
}

// WorkerReady 将准备好的worker推送到队列
func (s *Queue) WorkerReady(worker chan types.Request) {
	go func() {
		s.WorkerChan <- worker
	}()
}

// Run 任务调度中心
func (s *Queue) Run() {
	// 初始化request和worker队列通信通道
	s.RequestChan = make(chan types.Request)
	s.WorkerChan = make(chan chan types.Request)

	go func() {
		// 初始化队列
		var requestQueue []types.Request
		var workerQueue []chan types.Request
		for {
			// 判断空闲的request是否有woker进行接接收
			var activeRequest types.Request
			var activeWorker chan types.Request
			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}

			// select 进行任务调度
			select {
			case request := <-s.RequestChan:
				requestQueue = append(requestQueue, request)
			case worker := <-s.WorkerChan:
				workerQueue = append(workerQueue, worker)
			case activeWorker <- activeRequest:
				// 给请求分配对应的工作者，出队列
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]
			}
		}
	}()

}
