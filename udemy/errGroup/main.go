package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 800* time.Millisecond)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)
	s := []string{"task1", "task2", "task4", "task5", "task3"}
	for _, task := range s {
		task := task // avoid loopclosure
		eg.Go(func() error {
			return doTask(ctx, task)
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	
	fmt.Println("finished")
}

func doTask(ctx context.Context, task string) error {
	var t *time.Ticker

	switch task {
	case "task1":
		t = time.NewTicker(500 * time.Millisecond)
	case "task2":
		t = time.NewTicker(700 * time.Millisecond)
	default:
		t = time.NewTicker(1000 * time.Millisecond)
	}
	select {
	case <-ctx.Done():
		fmt.Printf("%v canceled : %v\n", task, ctx.Err())
		return ctx.Err()
	case <-t.C:
		t.Stop()
		// if task == "fake1" || task == "fake2" {
		// 	return fmt.Errorf("%v process failed", task)
		// }
		fmt.Printf("%v completed\n", task)
	}

	return nil
}
