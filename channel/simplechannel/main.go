package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		channel <- 10
	}()

	fmt.Println(<-channel)
}
