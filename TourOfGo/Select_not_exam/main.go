package main

import "fmt"

/*
fibonacci funciton.
selectは、複数のcaseのいずれかが準備できるようになるまでブロックし、準備ができたcaseを実行する。
この場合だとc, quitのどちらかが準備できるまでブロックし、準備ができた方を実行する。
*/

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	// c, quitをチャネルとして初期化。バッファはなし。
	c := make(chan int)
	quit := make(chan int)
	/*
		goroutineなので、非同期処理
		この関数の下にあるfibonacciが先に実行される
		iが10になるまでforループでcの値を受信する。つまり10回受信する。受信し終わったら、quitに0を送信する
	*/

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	/*
		最初に実行される
		関数内で、xの値をcに送信している。送信された時点でcの値が準備される
		cから値を取り出せるため、上記のgoroutineで受信できる
	*/
	fibonacci(c, quit)
}
