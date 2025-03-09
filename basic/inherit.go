package main

import "fmt"

func main6() {
	type User struct {
		Name string
		Age  int
	}

	// type Video struct {
	// 	Length int
	// 	Name   string
	// 	Author User
	// }
	type Video struct {
		Length int
		Name   string
		User   // 匿名フィールド
	}

	u := User{Name: "Kurosawa", Age: 20}
	// v := Video{Length: 100, Name: "go", Author: u}
	v := Video{Length: 100, Name: "go", User: u} // Userタイプを変数に格納
	fmt.Println(v.Length)
	fmt.Println(v.Age)       // VideoはUserのフィールドを継承している
	fmt.Println(v.User.Name) // Videoは匿名フィールドは直接アクセスできる

	// s := Student{1, 59, Residence{"Tokyo", "Shinjuku"}, nil}
	// or
	s := Student{1, 59, Residence{"Tokyo", "Shinjuku"}, &Student{}}

	// %vで変数の値を表示
	fmt.Printf("%v\n", s)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%#v\n", s)

	s2 := Student2{1, 58, struct {
		City    string
		Country string
	}{"Tokyo", "Japan"}, nil}
	fmt.Printf("%v\n", s2)
}

type Residence struct {
	Province string
	City     string
}

type Student struct {
	Id        int
	Score     float64
	Residence Residence
	Father    *Student // Student型のポインタ &Student{}
}

type TreeNode struct {
	Value      int
	LeftChild  *TreeNode
	RightChild *TreeNode
}

type Student2 struct {
	Id        int
	Score     float64
	Residence struct { // 匿名フィールド
		City    string
		Country string
	}
	Father *Student2
}
