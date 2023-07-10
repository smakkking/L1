package t25

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"time"
)

type TimeScheduler struct {
	Sign chan os.Signal
	Done chan bool
}

var T = TimeScheduler{
	Sign: make(chan os.Signal, 1),
	Done: make(chan bool, 1),
}

func (t *TimeScheduler) sleep(sec int) {

	wg := sync.WaitGroup{}

	wg.Add(1)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(sec)*time.Second)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}(ctx)
	wg.Wait()
}

func sleep(sec int) {
	var t time.Time = time.Now().Add(time.Duration(sec) * time.Second)

	for {
		if time.Now().Unix() >= t.Unix() {
			break
		}
	}
}

func main() {
	signal.Reset()
}
