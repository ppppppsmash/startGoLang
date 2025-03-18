package main

import "time"

func main() {
	c, d := 3, 5
	_, _ = c, d
	arg1(c, d)
}

func arg1(a int, b int) {
	a = a + b
	return
}

func arg2(a, b int) {
	a = a + b
}

// ポインタを渡す ポインタはメモリアドレスを指す
func arg3(a, b *int) {
	*a = *a + *b
	*b = 888
}

func return1(a, b int) int {
	a = a + b
	c := a
	return c
}

// cは返り値の変数名として事前に宣言されている
func return2(a, b int) (c int) {
	a = a + b
	c = a
	return
}

// 複数の返り値を返す
func return3() (int, int) {
	now := time.Now()
	return now.Hour(), now.Minute()
}
