package main

import (
	"fmt"
)

func main13() {
	array1d()
	array2d()
}

func array1d() {
	// 配列の宣言、値は0で初期化される
	var arr1 [5]int = [5]int{}
	var arr2 = [5]int{}
	var arr3 = [5]int{3, 2}
	// 2番目と4番目の値を15と30で初期化
	var arr4 = [5]int{2: 15, 4: 30}

	var arr5 = [...]int{3, 2, 6, 5, 4}
	var arr6 = [...]struct {
		name string
		age  int
	}{{"Tom", 18}, {"Jim", 20}}

	fmt.Printf("arr1=%#v\n", arr1)
	fmt.Printf("arr2=%#v\n", arr2)
	fmt.Printf("arr3=%#v\n", arr3)
	fmt.Printf("arr4=%#v\n", arr4)
	fmt.Printf("arr5=%#v\n", arr5)
	fmt.Printf("arr6=%#v\n", arr6)

	fmt.Printf("arr5[0]=%d\n", arr5[0])
	// 最後の要素を取得
	fmt.Printf("arr5[len(arr5)-1]=%d\n", arr5[len(arr5)-1])
	fmt.Printf("arrのアドレス=%p\n", &arr5)
	fmt.Printf("1個目のアドレス=%p\n", &arr5[0])
	fmt.Printf("2個目のアドレス=%p\n", &arr5[1])

	for i, ele := range arr5 {
		fmt.Printf("index=%d, element=%d\n", i, ele)
	}

	for i := 0; i < len(arr5); i++ {
		fmt.Printf("index=%d, element=%d\n", i, arr5[i])
	}

	fmt.Printf("len(arr1)=%d\n", len(arr1))
	fmt.Printf("len(arr5)=%d\n", cap(arr5))
}

func array2d() {
	// 5行3列の配列を宣言
	// 1行目は1だけ、2行目は2と3だけ
	// 列は空のまま
	var arr1 = [5][3]int{{1}, {2, 3}}

	var arr2 = [...][3]int{{1}, {2, 3}}

	fmt.Printf("arr[1][1]=%d\n", arr1[1][1])
	fmt.Printf("arr[4][2]=%d\n", arr1[4][2])

	for row, array := range arr2 {
		for col, ele := range array {
			fmt.Printf("arr2[%d][%d]=%d\n", row, col, ele)
		}
	}

	fmt.Printf("len(arr1)=%d\n", len(arr1))
	fmt.Printf("len(arr1)=%d\n", cap(arr1))
}

// 配列を渡す、オリジナルの配列を変更しない、コピーを渡すだけ
func update_array1(arr [5]int) {
	fmt.Printf("array in function, address is %p\n", &arr[0])
	arr[0] = 888
}

// ポインタを渡す、オリジナルの配列を変更する
func update_array2(arr *[5]int) {
	fmt.Printf("array in function, address is %p\n", &((*arr)[0]))
	arr[0] = 888
}

func update_array3(arr [5]*int) {
	*arr[0] = 111
}

func for_range_array() {
	arr := [...]int{1, 2, 3}

	// eleはコピーされる
	// コピーされたものを変更しても元の配列は変更されない
	for i, ele := range arr {
		arr[i] += 8
		fmt.Printf("%d %d %d\n", i, arr[i], ele)
		ele += 1
		fmt.Printf("%d %d %d\n", i, arr[i], ele)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d %d\n", i, arr[i])
	}
}
