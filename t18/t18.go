package t18

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	// закрываем мьютексом
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
}

func (c *Counter) Show() {
	fmt.Println(c.count)
}

func DoTask() {
	c := Counter{}
	defer c.Show()

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(c *Counter) {
			defer wg.Done()
			c.Increment()
		}(&c)
	}
	wg.Wait()
}
