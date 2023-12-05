package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 100; i <= 105; i++ {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 200; i <= 205; i++ {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 300; i <= 305; i++ {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	fmt.Println("Will executed after all go routines done")
}
