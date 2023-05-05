package main

import (
	"fmt"
	//"reflect"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

/*
TreeはそのLeftとRightにもTreeを持つ。
だから再帰的に呼び出すことで、すべてのnodeをwalkする。
再帰関数の終了条件は、nodeがない時。つまり、tがnilの時。
*/
func walk(t *tree.Tree, ch chan int) {
	// 終了条件
	if t == nil {
		return
	}
	/*
		Treeのnodeの値を、再帰的にチャネルに送信する。
		Left -> Value -> Rightの順なのは、二分木をソートされた順序で読み出すため
		この順序を変更してもコードは動くが、二分木の読み出し順が変わる
	*/
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// forループで、2つのチャネルを順に比較する
	for {
		// チャネルから値を受信
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		// ループの終了条件
		// どちらかのチャネルが閉じられたら、ループを抜ける
		if !ok1 || !ok2 {
			break
		}
		// 取り出された値が違う場合、その時点でfalseを返す
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	// 1から10までの値を持つ二分木を作成
	go Walk(tree.New(1), ch)
	// チャネルから値を受信し、出力する
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
