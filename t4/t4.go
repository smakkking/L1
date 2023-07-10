package t4

import (
	"context"
	"fmt"
)

func t4_f1() {

}

func DoTask() {
	ch := make(chan int)

	ctx, _ := context.WithCancel(context.Background())

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

}