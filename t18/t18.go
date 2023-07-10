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
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
}

func (c *Counter) Show() {
	fmt.Println(c.count)
}

func DoTask() {

}
