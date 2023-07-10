package t3

import (
	"fmt"
	"sync"
)

func t3_f1(arr []int) {
	// пишем горутинами в канал, затем суммируем в мейн-горутине
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

	var sum int = 0
	for elem := range ch {
		sum += elem
	}

	fmt.Println(sum)
}

func t3_f2(arr []int) {
	// пишем горутинами в канал, затем суммируем в отдельной горутине
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
		var sum int = 0
		for {
			select {
			case x := <-ch:
				sum += x
			default:
				mu.Lock()
				if n == len(arr) {
					mu.Unlock()
					fmt.Println(sum)
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

	t3_f2(arr)

	// stdout goroutine safe?
	// https://stackoverflow.com/questions/14694088/is-it-safe-for-more-than-one-goroutine-to-print-to-stdout

	// how goroutine takes var from outside area of ​​visibility

	var a int

	{
		a := "sdtst"
		fmt.Println(a)
	}

	fmt.Println(a)
}
