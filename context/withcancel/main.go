package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parentContext := context.Background()
	ctx, cancel := context.WithCancel(parentContext)
	defer cancel() 

	go performTask(ctx)

	time.Sleep(2 * time.Second)

	cancel()

	time.Sleep(2 * time.Second)
}

func performTask(ctx context.Context) {
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Task in progress...")
		case <-ctx.Done():
			fmt.Println("Task canceled.")
			return
		}
	}
}
