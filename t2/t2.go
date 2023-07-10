package t2

import (
	"fmt"
	"sync"
)

func t2_f1(arr []int) {
	wg := sync.WaitGroup{} // будем регистрировать все горутины

	for elem := range arr {
		// пихаем горутину в ВГ
		wg.Add(1)
		go func(i int) {
			defer wg.Done() // когда функция закончит работу, счетчик вг уменьшится
			fmt.Println(i * i)
		}(elem)
	}

	wg.Wait() // ждем, пока все запущеные горутины закончатся
}

func t2_f2(arr []int) {
	// здесь мы пишем квадраты в канал, а потом из него читаем
	ch := make(chan int, len(arr)) // буфф канал

	wg := sync.WaitGroup{}

	for elem := range arr {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- i * i // пишем в канал
		}(elem)
	}

	wg.Wait()
	close(ch)

	for elem := range ch {
		fmt.Println(elem)
	}
}

func t2_f3(arr []int) {
	// здесь мы горутинами пишем в канал, а потом отдельной горутиной из него читаем
	// (типо имитация, что приходит в канал извне и отдельная горутина ее обрабатывает)
	ch := make(chan int, len(arr))
	mu := sync.Mutex{}
	var n int = 0

	for elem := range arr {
		go func(i int) {
			ch <- i * i // пишем в канал

			mu.Lock()
			n++
			mu.Unlock()
		}(elem)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(ch chan int) {
		defer wg.Done()
		for {
			select {
			case x := <-ch:
				fmt.Println(x)
			default:
				mu.Lock()
				if n == len(arr) {
					mu.Unlock()
					return
				}
				mu.Unlock()
			}
		}
	}(ch)

	wg.Wait()
}

func DoTask() {
	arr := []int{2, 4, 6, 8, 10}

	t2_f3(arr)

	// stdout goroutine safe?
	// https://stackoverflow.com/questions/14694088/is-it-safe-for-more-than-one-goroutine-to-print-to-stdout

	// how goroutine takes var from outside area of ​​visibility
}
