package main

import (
	"fmt"
)

func main() {
	channel := make(chan string, 2)

	go func() {
		channel <- "im first channel"
		channel <- "im second channel"
	}()

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
