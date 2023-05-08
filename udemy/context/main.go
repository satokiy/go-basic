package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
  WithTimeoutによって複数のごルーチンを制御する
*/
// func main() {
// 	var wg sync.WaitGroup
// 	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
// 	defer cancel()
// 	wg.Add(3)
// 	go subTask(ctx, &wg, "A")
// 	go subTask(ctx, &wg, "B")
// 	go subTask(ctx, &wg, "C")
// 	wg.Wait()
// }

// func subTask(ctx context.Context, wg *sync.WaitGroup, id string) {
// 	defer wg.Done()
// 	// 500ms周期でidを出力する
// 	ticker := time.NewTicker(500 * time.Millisecond)
// 	defer ticker.Stop()
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println(ctx.Err())
// 			return
// 		case <-ticker.C:
// 			ticker.Stop()
// 			fmt.Println(id)
// 			return
// 		}
// 	}
// }

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	wg.Add(2)
	go func() {
		defer wg.Done()
		v, err := criticalTask(ctx)
		if err != nil {
			fmt.Println("critical task failed")
			cancel()
			return
		}
		fmt.Println("critical task success:", v)
	}()

	go func() {
		defer wg.Done()
		v, err := normalTask(ctx)
		if err != nil {
			fmt.Println("normal task failed")
			return
		}
		fmt.Println("normal task success:", v)
	}()

	wg.Wait()
	fmt.Println("main function exit")
}

func criticalTask(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 800*time.Millisecond)
	defer cancel()
	t := time.NewTicker(1000 * time.Millisecond)
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-t.C:
		t.Stop()
		fmt.Println("critical task exec")
	}
	return "A", nil
}

func normalTask(ctx context.Context) (string, error) {
	t := time.NewTicker(3000 * time.Millisecond)
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-t.C:
		t.Stop()
		fmt.Println("normal task exec")
	}
	return "B", nil
}
