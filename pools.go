package GoblinPools

import (
	"fmt"
	"sync"
)

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(numWorkers int) *Pool {
	pool := Pool{
		work: make(chan Worker),
	}

	pool.wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(id int) {

			for work := range pool.work {
				work.Task()
				fmt.Println("i am ", id)
			}
			pool.wg.Done()
		}(i)
	}

	return &pool
}

func (pool *Pool) Run(work Worker) {
	pool.work <- work
}

func (pool *Pool) Shutdown() {
	close(pool.work)
	pool.wg.Wait()
}
