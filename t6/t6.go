package t6

// https://stackoverflow.com/questions/6807590/how-to-stop-a-goroutine
// https://golang-blog.blogspot.com/2020/03/stop-goroutine.html

import (
	"context"
	"fmt"
	"sync"
)

func simple_goroutine1(ctx context.Context) {
	// горутинка, бегает в вечном цикле
	defer fmt.Println("I am stopped 1")

	for {
		select {
		case <-ctx.Done():
			break
		}
	}
}

func simple_goroutine2(stop chan bool) {
	// горутинка, бегает в вечном цикле
	defer fmt.Println("I am stopped 2")

	for {
		select {
		case x := <-stop:
			if x {
				break
			}
		default:
			// do some stuff
		}
	}
}

func simple_goroutine3(stop chan bool) {
	defer fmt.Println("I am stopped 3")

	for {
		_, ok := <-stop
		if !ok {
			break
		}
	}
}

func simple_goroutine4(m sync.Map) {
	defer fmt.Println("I am stopped 4")

	for {
		_, ok := m.Load("end")
		if ok {
			break
		}
	}
}

func DoTask() {
	ctx, cancel := context.WithCancel(context.Background())

	go simple_goroutine1(ctx)

	cancel()

	stop := make(chan bool)
	go simple_goroutine2(stop)

	stop <- true

	stop1 := make(chan bool)
	go simple_goroutine3(stop1)
	close(stop1)

}
