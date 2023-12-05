package main

import (
	"fmt"
	"sync"
)

var sharedSlice []int
var mutex sync.Mutex

func appendSlice(value int, wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock()
	sharedSlice = append(sharedSlice, value)
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go appendSlice(i, &wg)
	}

	wg.Wait()

	fmt.Println("Final Slice:", sharedSlice)
}
