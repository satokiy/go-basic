package main

import (
	"errors"
	"fmt"
)

type Hoge struct {
	Name string
}

func (h *Hoge) Error() string {
	return "hoge error"
}

func main() {
	err01 := errors.New("error 01")
	fmt.Printf("%[1]p %[1]T %[1]v\n", err01)

	err02 := errors.New("error 02")
	fmt.Printf("%[1]p %[1]T %[1]v\n", err02)
}
