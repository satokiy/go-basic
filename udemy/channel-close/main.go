package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch1)
	}()

	ch1 <- 1
	close(ch1)

	v, ok := <-ch1
	fmt.Println(v, ok) // 0 false
	wg.Wait()

	ch2 := make(chan int, 2)
	ch2 <- 10
	ch2 <- 20
	close(ch2)

	v2, ok2 := <-ch2
	fmt.Println(v2, ok2) // 10 true
	v3, ok3 := <-ch2
	fmt.Println(v3, ok3) // 20 true
	v4, ok4 := <-ch2
	fmt.Println(v4, ok4) // 0 false
}

func generateCountStream() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	return ch
}
