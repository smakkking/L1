package t5

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func writer(ctx context.Context, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	var i = 0
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(100 * time.Millisecond)
			// если канал небуф - то будет deadlock, потому что из читателя мы уже выйдем,
			// и попытаемся записать в небуф канал без читателя и заблокируемся навечно, главная горутина тоже стоит и ждет завершения этой
			ch <- i
			i++
		}

	}
}

func reader(ctx context.Context, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case x := <-ch:
			fmt.Println(x)
		}
	}
}

func DoTask() {
	my_ch := make(chan int, 1) // можно и с небуф каналом
	wg := sync.WaitGroup{}

	var N int
	fmt.Printf("F=")
	fmt.Scanf("%d", &N)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(N)*time.Second)

	wg.Add(1)
	go writer(ctx, my_ch, &wg)

	wg.Add(1)
	go reader(ctx, my_ch, &wg)

	wg.Wait()
}
