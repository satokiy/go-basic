package main

import (
	"fmt"
	"runtime"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println(<-ch)
	}()
	ch <- 10
	fmt.Printf("number of goroutine: %v\n", runtime.NumGoroutine())

	ch2 := make(chan int, 1)
	ch2 <- 2
	fmt.Println(<-ch2)

}
