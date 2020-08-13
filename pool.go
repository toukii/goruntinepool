package goruntinepool

import (
	"fmt"
	"sync"
)

type Pool struct {
	n     int
	wg    sync.WaitGroup
	tasks chan *task
}

type task struct {
	handler func() error
}

func NewPool(n int) *Pool {
	pool := &Pool{
		n:     n,
		tasks: make(chan *task, n<<1),
	}
	pool.start()

	return pool
}

func (p *Pool) Run(fn func() error) {
	p.wg.Add(1)
	p.tasks <- &task{
		handler: fn,
	}
}

func (p *Pool) start() {
	for i := 0; i < p.n; i++ {
		idx := i
		go p.goruntine(idx)
	}
}

func (p *Pool) goruntine(i int) {
	for {
		fmt.Printf("## goruntine %d\n", i)
		tk, ok := <-p.tasks
		if !ok { // close
			return
		}
		func() {
			defer p.wg.Done()
			err := tk.handler()
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
}

func (p *Pool) Wait() {
	p.wg.Wait()
	close(p.tasks)
}
