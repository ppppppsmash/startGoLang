package main

import "fmt"

func main() {
	if 5 > 9 {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}

	// 変数を宣言してからif文を書く
	if score := 60; score > 60 {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}

	// 複数の変数を宣言してからif文を書く
	if c, d, e := 1, 4, 7; c > d {
		fmt.Println("OK")
	} else if c > e {
		fmt.Println("OK2")
	} else {
		fmt.Println("NG")
	}

	if f, g, h := 1, 4, 7; f < g && (f > 0 || h > 3) {
		fmt.Println("OK")
	} else if f > g {
		fmt.Println("OK2")
	} else {
		fmt.Println("NG")
	}

}
