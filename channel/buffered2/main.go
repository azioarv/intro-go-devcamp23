package main

import (
	"fmt"
)

func main() {
	bufferedChan := make(chan int, 5)

	go func() {
		for i := 1; i <= 5; i++ {
			bufferedChan <- i
			fmt.Printf("Sent: %d\n", i)
		}
		close(bufferedChan)
	}()

	for num := range bufferedChan {
		fmt.Printf("Received: %d\n", num)
	}
}