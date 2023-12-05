package main

import (
	"context"
	"fmt"
	"time"
)

func performTask(ctx context.Context) error {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Task completed successfully")
		return nil
	case <-ctx.Done():
		fmt.Println("Task canceled due to deadline exceeded")
		return ctx.Err()

	}
}

func main() {
	deadline := time.Now().Add(6 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	err := performTask(ctx)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
