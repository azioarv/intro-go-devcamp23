package main

import (
	"fmt"
	"time"
)

func main() {
	go async()
	for counter := 200; counter <= 205; counter++ {
		fmt.Println(counter)
		time.Sleep(100 * time.Millisecond)
	}
}

func async() {
	for counter := 100; counter <= 105; counter++ {
		fmt.Println(counter)
		time.Sleep(100 * time.Millisecond)
	}
}
