package main

import "fmt"

func main4() {
	const PI float32 = 3.14

	const (
		URL      = "http://example.com"
		SiteName = "Example"
	)

	const (
		a = iota // 0
		b        // 1
		c        // 2
		d        // 3
	)

	const (
		// ビット計算
		// KB = 1 << 10
		// MB = 1 << 20
		// GB = 1 << 30

		// ビット計算 進化1
		// NOT_USE = iota // iota=0
		// KB = 1 << (iota * 10)
		// MB = 1 << (iota * 10)
		// GB = 1 << (iota * 10)

		// ビット計算 進化2
		NOT_USE = 1 << (iota * 10) // iota=0
		KB                         // iota=1
		MG                         // iota=2
		GB                         // iota=3
	)

	const (
		ss, mm = iota + 1, iota + 2 // iota=0
		bb, cc                      // iota=1
		dd, ee                      // iota=2
	)

	// {}ブロックスコープ
	{
		const PI float32 = 3.14
		fmt.Printf("PI=%f\n", PI)
	}
}
