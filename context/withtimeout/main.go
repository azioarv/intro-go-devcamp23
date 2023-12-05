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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := performTask(ctx)
	if err != nil {
		fmt.Println("Error:", err)
	}
}