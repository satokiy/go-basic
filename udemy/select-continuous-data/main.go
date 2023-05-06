package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const bufSize = 5

func main() {
	ch1 := make(chan int, bufSize)
	ch2 := make(chan int, bufSize)
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Millisecond)

	defer cancel()
	wg.Add(3)
	go countProducer(&wg, ch1, bufSize, 50)
	go countProducer(&wg, ch2, bufSize, 500)
	go countConsumer(ctx, &wg, ch1, ch2)
	wg.Wait()
	fmt.Println("finish")

}

/*
size の数だけチャネルに値を格納する
*/
func countProducer(wg *sync.WaitGroup, ch chan<- int, size int, sleep int) {
	defer wg.Done()
	defer close(ch)

	for i := 0; i < size; i++ {
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		ch <- i
	}
}

/*
引数のチャネルから値を受信して表示する
*/
func countConsumer(ctx context.Context, wg *sync.WaitGroup, ch1 <-chan int, ch2 <-chan int) {
	defer wg.Done()

	/*
		引数に渡された2つのチャネルの値を、どちらかを受信するまでループ
		コンテキストが終了した場合、それを出力
	*/
	for ch1 != nil || ch2 != nil {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			for ch1 != nil || ch2 != nil {
				select {
				case v, ok := <-ch1: // caseの変数のスコープはcase内に閉じる
					if !ok {
						ch1 = nil
						break
					}
					fmt.Printf("[ch1] %v\n", v)
				case v, ok := <-ch2:
					if !ok {
						ch2 = nil
						break
					}
					ch2 = nil
					fmt.Printf("[ch2] %v\n", v)
				}
			}
		case v, ok := <-ch1: // caseの変数のスコープはcase内に閉じる
			if !ok {
				ch1 = nil
				break
			}
			fmt.Printf("[ch1] %v\n", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			ch2 = nil
			fmt.Printf("[ch2] %v\n", v)
		}
	}
}
