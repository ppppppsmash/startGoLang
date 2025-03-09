package main

import "fmt"

type User struct {
	Id            int
	Score         float64
	Name, Address string
}

func main() {
	var u User
	u = User{Id: 1, Score: 100, Name: "John", Address: "Tokyo"}
	// u = User{1, 100, "John", "Tokyo"} // 順番を変えるとエラーになる
	fmt.Println(u.Name)
	u.hello()

	// type student struct {
	// 	Name string
	// 	Age  int
	// }

	// 匿名構造体 一時的に使う
	var student struct {
		Name string
		Age  int
	}
	student.Name = "John"
	student.Age = 20
	fmt.Println(student)

	u2 := User{}
	// var u2 User
	u2.Address = "Osaka"

	// ポインタ 参照渡し メモリアドレスを渡す
	u3 := &u2
	fmt.Println(u3)

	// new() メモリアドレスを返す
	u4 := new(User)
	fmt.Println(u4)

}

// メソッド
func (u User) hello() {
	fmt.Println("Hello, my name is", u.Name)
}
