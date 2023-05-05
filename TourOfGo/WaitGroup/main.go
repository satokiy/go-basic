package main

import (
	"fmt"
	"sync"
	// "runtime"
	"time"
)

func getMinutesAndSeconds() string {
	return time.Now().Format("04:05")
}

func test(n int, wg *sync.WaitGroup) {
	fmt.Printf("[test start]\t%v\t[%v]\n", n, getMinutesAndSeconds())
	time.Sleep(5*time.Second)
	fmt.Printf("[test end]\t%v\t[%v]\n", n, getMinutesAndSeconds())
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	s := time.Now()
	fmt.Printf("[main start]\t[%v]\n", getMinutesAndSeconds())
	
	for i:=0; i<3; i++ {
		wg.Add(1)
		go test(i, &wg)
	}
	// test(1)
	// fmt.Printf("NumCPU: %v\n", runtime.NumCPU())
	// fmt.Printf("NumGoroutine: %v\n", runtime.NumGoroutine())
	wg.Wait()
	fmt.Printf("[main end]\t[%v]\n", getMinutesAndSeconds())
	e := time.Now()
	fmt.Printf("time: %v\n", e.Sub(s).Round(time.Second))
}