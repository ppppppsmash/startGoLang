package main

import "fmt"

func main11() {
	var a uint64 = 5        // 5 = 4 + 1
	fmt.Printf("a=%b\n", a) // 二進数
	binaryFormat(a)
}

func binaryFormat(a uint64) {
	var c uint64 = 1 << 63
	// if c&a == c {
	// 	fmt.Print("1")
	// } else {
	// 	fmt.Print("0")
	// }

	// c = c >> 1
	// if c&a == c {
	// 	fmt.Print("1")
	// } else {
	// 	fmt.Print("0")
	// }

	for i := 0; i < 64; i++ {
		if c&a == c {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
		c = c >> 1
	}
}
