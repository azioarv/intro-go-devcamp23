package main

import (
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	limiter := rate.NewLimiter(4, 4)

	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Printf("Action %d: Performed at %s\n", i+1, time.Now().Format(time.StampMilli))
		} else {
			fmt.Println("Too many request")
		}
		time.Sleep(50 * time.Millisecond)
	}
}
