package t4

import (
	"context"
	"fmt"
)

func DoTask() {
	ch := make(chan int, 1)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const N int = 10

	for i := 0; i < N; i++ {
		go func(ctx context.Context, ch <-chan int) {
			for {
				select {
				case x := <-ch:
					fmt.Println(x)
				case <-ctx.Done():
					return
				}
			}
		}(ctx, ch)
	}

	for i := 0; true; i++ {
		ch <- i
	}

}
