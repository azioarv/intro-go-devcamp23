package main

import (
	"fmt"
	"time"
)

func main() {
	throttleInterval := 1 * time.Second
	ticker := time.NewTicker(throttleInterval)
	defer ticker.Stop()

	for i := 0; i < 5; i++ {
		<-ticker.C
		fmt.Printf("Action %d: Performed at %s\n", i+1, time.Now().Format(time.StampMilli))
	}
}
