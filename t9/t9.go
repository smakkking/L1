package t9

import (
	"fmt"
	"sync"
	"time"
)

// записывает числа в канал
func x1(arr []int, in chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, elem := range arr {
		in <- elem
		time.Sleep(500 * time.Millisecond)
	}
	close(in)
}

// читает числа из канала и записывает их * 2 в другой
func transfer(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for elem := range in {
		out <- 2 * elem
	}
	close(out)
}

// читает из канала и выводит
func x2(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for elem := range in {
		fmt.Println(elem)
	}
}

func DoTask() {
	ch_in := make(chan int)
	ch_out := make(chan int)

	wg := sync.WaitGroup{}

	var arr = []int{1, 2, 3, 4, 5}

	wg.Add(1)
	go x1(arr, ch_in, &wg)

	wg.Add(1)
	go transfer(ch_in, ch_out, &wg)

	wg.Add(1)
	go x2(ch_out, &wg)

	wg.Wait()
}
