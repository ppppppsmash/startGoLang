package main

import "fmt"

// Excelの最後の列はXFD、じゃあExcelの総列は？
// 26進数で考える
// A: 1, B: 2, C: 3, ..., Z: 26
// AA: 27, AB: 28, ..., AZ: 52, BA: 53, BB: 54, ..., ZZ: 701
// AAA: 702, AAB: 703, ..., ZZZ: 18278
// 26進数で考える
// A: 1, B: 2, C: 3, ..., Z: 26
// AA: 27, AB: 28, ..., AZ: 52, BA: 53, BB: 54, ..., ZZ: 701
// AAA: 702, AAB: 703, ..., ZZZ: 18278

func main() {
	fmt.Println("A=%d, Z=%d\n", 'A', 'Z')

	var base int = 'Z' - 'A' + 1
	fmt.Println("base", base)

	var total int
	total += 'D' - 'A' + 1
	total += base * ('F' - 'A' + 1)
	total += base * base * ('X' - 'A' + 1)
	fmt.Println("total", total)
}
