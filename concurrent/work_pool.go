package concurrent

import (
	"sync"
)

type Worker func()

type WorkerPool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func NewWorkerPool(routineNum int, chanSize int) *WorkerPool {
	p := &WorkerPool{
		work: make(chan Worker, chanSize),
	}

	p.wg.Add(routineNum)
	for i := 0; i < routineNum; i++ {
		go func() {
			for w := range p.work {
				w()
			}
			p.wg.Done()
		}()
	}
	return p
}

func (p *WorkerPool) Run(w Worker) {
	p.work <- w
}

func (p *WorkerPool) ShutDown() {
	close(p.work)
	p.wg.Wait()
}
