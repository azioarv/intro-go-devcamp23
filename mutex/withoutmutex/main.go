package main

import (
	"fmt"
	"sync"
)

var sharedSlice []int
var wg sync.WaitGroup

func appendSlice(value int) {
	defer wg.Done()
	sharedSlice = append(sharedSlice, value)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go appendSlice(i)
	}

	wg.Wait()

	fmt.Println("Final Slice:", sharedSlice)
}
