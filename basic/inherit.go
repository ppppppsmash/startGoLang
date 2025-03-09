package main

import "fmt"

func main() {
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
}
