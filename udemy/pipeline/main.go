package main

import (
	"context"
	"fmt"
)

// func: generator - return a readonly channel
// args: context, ...int
// return: <-chan int
func generator(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()
	return out
}

// func: double
// args: context, <-chan int
// return: chan int
func double(ctx context.Context, in <-chan int) (out chan int) {
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-ctx.Done():
				return
			case out <- n * 2:
			}
		}
	}()
	return
}

func offset(ctx context.Context, in <-chan int) (out chan int) {
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-ctx.Done():
				return
			case out <- n + 2:
			}
		}
	}()
	return
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nums := []int{1, 2, 3, 4, 5}
	var i int
	flag := true

	for v := range double(ctx, offset(ctx, double(ctx, generator(ctx, nums...)))) {
		if i == 3 {
			cancel()
			flag = false
		}
		if flag {
			fmt.Println(v)
		}
		i++
	}

	fmt.Println("finish")
}
