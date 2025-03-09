package main

import "fmt"

func main8() {
	var sum int // sumは0
	for a := 6; a >= 0; a-- {
		sum += a
	}
	fmt.Println(sum)

	// for c, d := 6, 30; c >= 0; c, d = c-1, d/2 {
	// 	fmt.Println(c, d)
	// }
	for c, d := 6, 30; c >= 0; {
		fmt.Println(c, d)
		c, d = c-1, d/2
	}

	e, f := 8, 40
	for e >= 0 && f > e {
		sum += e
		if e == 5 {
			break // breakでループを抜ける
		}
		e, f = e-1, f/2

		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if j == 2 {
					continue // continueでループをスキップ
				}
				fmt.Println(i, j)
			}
		}
	}
	fmt.Println(sum, e, f)
}
