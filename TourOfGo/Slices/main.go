package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) (pic [][]uint8) {
	pic = make([][]uint8, dy)
	for i := range pic {
		pic[i] = make([]uint8, dx)
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
