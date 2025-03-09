package main

import "fmt"

func main() {
	// basic_goto()
	// if_goto()
	// for_goto()
	// continue_label()
	break_label()
}

func basic_goto() {
	var i int = 4

	// abc: ラベル gotoで使う
MY_LABEL:
	i += 3
	i *= 2
	fmt.Println(i)

	if i > 30 {
		return
	}

	// if分がないとMY_LABELは無限ループになる
	goto MY_LABEL
}

func if_goto() {
	var i int = 10
	if i%2 == 0 {
		goto L1 // 偶数の場合、L1にジャンプ
	} else {
		goto L2 // 奇数の場合、L2にジャンプ
	}

L1:
	i += 3
	fmt.Println(i)
L2:
	i *= 3
	fmt.Println(i)
}

func for_goto() {
	const SIZE = 5

L1:
	for i := 0; i < SIZE; i++ {
	L2:
		fmt.Printf("第%d行目をチェック", i)
	L3:
		if i%2 == 1 {
			for j := 0; j < SIZE; j++ {
				fmt.Printf("第%d列目をチェック", j)

				if j%3 == 0 {
					goto L1
				} else if j%3 == 1 {
					goto L2
				} else {
					goto L3
				}
			}
		}
	}
}

func continue_label() {
	const SIZE = 5
L1:
	for i := 0; i < SIZE; i++ {
	L2:
		fmt.Printf("第%d行目をチェック\n", i)
	L3:
		if i%2 == 1 {
			for j := 0; j < SIZE; j++ {
				fmt.Printf("第%d列目をチェック\n", j)

				if j%3 == 0 {
					// continueの場合、labelは絶対に外側(for分)ループになる
					continue L1
				} else if j%3 == 1 {
					goto L2
				} else {
					goto L3
				}
			}
		}
	}
}

func break_label() {
	const SIZE = 5
L1:
	for i := 0; i < SIZE; i++ {
	L2:
		fmt.Printf("第%d行目をチェック\n", i)

		if i%2 == 1 {
		L3:
			for j := 0; j < SIZE; j++ {
				fmt.Printf("第%d列目をチェック\n", j)

				if j%3 == 0 {
					// breakの場合、labelは絶対に外側(for分)ループになる
					break L1
				} else if j%3 == 1 {
					goto L2
				} else {
					goto L3
				}
			}
		}

	}
}
