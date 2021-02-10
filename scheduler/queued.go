package scheduler

import "douban-book-crawler/engine"

/*
维护一个 Request 队列和一个 Worker 队列
有 Request 就加入 Request 队列
有 Worker 就加入 Worker 队列
两个队列都不为空的时候就把 Request 发给 Worker
*/

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueuedScheduler) WorkerReady(w chan engine.Request) {
	q.workerChan <- w
}

func (q *QueuedScheduler) ConfigureMasterWorkerChan(r chan engine.Request) {
	q.workerChan <- r
}

func (q *QueuedScheduler) Run() {
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			// 两个队列都不为空的时候，取出队首元素，同时保证了在 select 里面出队
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			// 加入 Request 队列
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			// 加入 Worker 队列
			case w := <-q.workerChan:
				workerQ = append(workerQ, w)
			// 出队
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
