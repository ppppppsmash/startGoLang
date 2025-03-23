package main

import (
	"fmt"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func main() {
	slice_init()
	expansion()
}

func slice_init() {
	// len=cap=0
	var s []int
	fmt.Printf("len %d cap %d \n", len(s), cap(s))
	// len=cap=0
	s = []int{}
	fmt.Printf("len %d cap %d \n", len(s), cap(s))
	// len=cap=3
	s = make([]int, 3)
	fmt.Printf("len %d cap %d \n", len(s), cap(s))

	// len=3, cap=5
	s = make([]int, 3, 5)
	fmt.Printf("len %d cap %d \n", len(s), cap(s))

	// len=cap=5
	s = []int{1, 2, 3, 4, 5}
	fmt.Printf("len %d cap %d \n", len(s), cap(s))
	fmt.Println("===================")

	s2d := [][]int{
		{1},
		{2, 3},
	}
	fmt.Printf("s2d len=%d cap=%d\n", len(s2d), cap(s2d))
	fmt.Printf("s2d[0] len=%d cap=%d\n", len(s2d[0]), cap(s2d[0]))
	fmt.Printf("s2d[1] len=%d cap=%d\n", len(s2d[1]), cap(s2d[1]))
}

func expansion() {
	s := make([]int, 0, 3)
	prevCap := cap(s)
	for i := 0; i < 100; i++ {
		s = append(s, i)
		fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
		currCap := cap(s)
		if currCap > prevCap {
			fmt.Printf("capacityは%dから%dに増えました\n", prevCap, currCap)
			prevCap = currCap
		}
	}
}
