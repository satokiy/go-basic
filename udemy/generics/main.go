package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type customConstraints interface {
	string | ~int | float64 | float32
}

type NewInt int

func add[T customConstraints](x, y T) T {
	return x + y
}

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(min(3, 4))
	m2 := map[int]float32{
		1: 1.1,
		2: 2.2,
	}
	fmt.Println(min(m2[1], m2[2]))

}
