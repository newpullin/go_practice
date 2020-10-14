package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type counter1 struct {
	i int64
}

func (c *counter1) increment() {
	atomic.AddInt64(&c.i, 1)
}

func (c *counter1) display() {
	fmt.Println(c.i)
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter1{i: 0}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.increment()
		}()
	}

	wg.Wait()

	c.display()
}
