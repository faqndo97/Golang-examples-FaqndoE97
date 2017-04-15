package main

import (
	"fmt"
	"sync"
)

func meter(val int, ch chan<- int) {
	select {
	case ch <- val:
		fmt.Println("a")
	default:
		fmt.Println("b")
	}
}

func sacar(ch <-chan int) {
	select {
	case val := <-ch:
		fmt.Println(val)
	default:
		fmt.Println("c")
	}
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)
	wg.Add(3)
	go func(ch chan<- int) {
		meter(1, ch)
		wg.Done()
	}(ch)
	for i := 1; i <= 2; i++ {
		go func(ch <-chan int) {
			go sacar(ch)
			wg.Done()
		}(ch)
	}
	wg.Wait()
}
