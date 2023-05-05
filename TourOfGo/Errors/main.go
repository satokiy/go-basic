package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		z := 1.0
		diff := float64(10000)         // 十分に大きな値
		stop_diff := math.Pow(10, -10) // loopを止める値
		for math.Abs(diff) > stop_diff {
			diff = (z*z - x) / (2 * z)
			z -= diff

		}
		return z, nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
