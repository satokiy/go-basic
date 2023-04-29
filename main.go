package main

import (
	"fmt"
	"unsafe"
)

const secret = "abc"

type Os int

const (
	Mac Os = iota + 1
	Windows
	Linux
)


func main() {
	var ui1 uint16
	// var ui2 uint16
	fmt.Printf("memory address of ui1: %p\n", &ui1)
	// fmt.Printf("memory address of ui2: %p\n", &ui2)
	
	var p1* uint16
	fmt.Printf("value of p1: %v\n", p1)
	p1 = &ui1
	fmt.Printf("value of p1: %v\n", p1)
	fmt.Printf("size of p1: %d[byte]\n", unsafe.Sizeof(p1))
	fmt.Printf("memory address of p1: %p\n", &p1)
	*p1 = 10
	fmt.Printf("value of ui1 dereference: %v\n", *p1)
	fmt.Printf("value of ui1: %v\n", ui1)

	var pp1** uint16 = &p1
	fmt.Printf("value of pp1: %v\n", pp1)

	// var p2* uint16
	// p2 = &ui2
	// fmt.Printf("value of p2: %v\n", p2)
	
	

}
