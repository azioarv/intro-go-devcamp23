package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Producing:", i)
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel closed. Exiting consumer.")
			return
		}
		fmt.Println("Consumed:", val)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)
	go func() {
		defer wg.Done()
		producer(ch)
	}()

	go func() {
		defer wg.Done()
		consumer(ch)
	}()
	wg.Wait()
}
