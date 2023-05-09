package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	cores := runtime.NumCPU()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	outChs := make([]<-chan string, cores)
	inData := generator(ctx, nums...)

	fanOut(ctx, generator(ctx, cores))

}

// args: ...int
// return: chan int
func generator(ctx context.Context, nums ...int) (out chan int) {
	go func() {
		defer close(out)
		for _, num := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- num:
			}
		}
	}()
	return
}

func fanOut(ctx context.Context, in <-chan int) <-chan string {
	out := make(chan string)
	go func() {
		// 重たい処理
		heavyWork := func(i int, id int, sleepTime time.Duration) string {
			time.Sleep(sleepTime)
			return fmt.Sprintf("result: %v (id: %v)", i*i, id)
		}
		for v := range in {
			select {
			case <-ctx.Done():
				return
			case out <- heavyWork(v, 1, 200*time.Microsecond):
			}
		}
	}()
	return out
}

func fanIn(ctx context.Context, chs ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	
	multiplex := func(ch <-chan string) {
		defer wg.Done()
		for v := range ch {
			select {
			case <-ctx.Done():
				return
			case out <- v:
			}
		}
	}
	wg.Add(len(chs))
	for _, ch := range chs {
		go multiplex(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
