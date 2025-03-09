package main

import "fmt"

func main2() {
	var MyName string = "Kurosawa"
	fmt.Println(MyName)

	var a int = 8
	var b = a

	// _ は未使用の変数を表す
	_ = b

	// := は変数の宣言と代入を同時に行う  : var を省略できる
	c := b // var c = b と同じ

	// 変数の値を変更する
	a = c

	// var() は複数の変数を宣言する
	var (
		d uint16
		e int8
		f float32
		g float64
	)

	a = -5

	// 05は8進数
	d = 05
	// 0o57は8進数
	a = 0o57

	// 0x12は16進数
	a = 0x12

	// 5_0_123_7は10進数
	a = 5_0_123_7

	// bool型はtrueかfalseのみ
	var n bool = true

	// 変数を宣言して未使用の場合は _ を使用する
	_, _, _, _ = d, e, f, g

	fmt.Println(n)

	// %d は整数, %f は浮動小数点数, %t はbool
	fmt.Printf("d=%d, e=%f, f=%t, g=%f\n", d, e, f, g)

	// %f は浮動小数点数, %g は浮動小数点数の省略, %e は浮動小数点数の指数表示 [1] は1番目の引数を参照 [2] は2番目の引数を参照
	fmt.Printf("f=%[1]f, g=%[2]f, g=%[2]g, f=%[1]e\n", f, g)

}
