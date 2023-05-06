package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("goroutine!!")
		defer wg.Done()
	}()

	wg.Wait()
	fmt.Printf("num of working goroutine: %v\n", runtime.NumGoroutine())
	fmt.Println("main function!!")

}
