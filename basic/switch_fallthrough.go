package main

import "fmt"

func main() {
	// basic_switch()
	// switch_expression()
	fall_through(60)
}

func basic_switch() {
	color := "red"

	switch color {
	case "red":
		fmt.Println("赤色です")
	case "blue":
		fmt.Println("青色です")
	case "green":
		fmt.Println("緑色です")
	default:
		fmt.Println("色がありません")
	}
}

func add(a int) int {
	return a + 10
}

func switch_expression() {
	var a int = 5
	switch add(a) {
	case 15:
		fmt.Println("right")
	default:
		fmt.Println("wrong")
	}

	// 式なし
	switch {
	case add(a) == 15:
		fmt.Println("right")
	default:
		fmt.Println("wrong")
	}
}

// fallthroughはcaseの中でのみ有効、条件に満たすとその下のcaseも実行される
func fall_through(age int) {
	fmt.Printf("age=%d\n", age)
	switch {
	case age > 70:
		fmt.Println("老人です")
		fallthrough
	case age > 30:
		fmt.Println("中年です")
		fallthrough
	case age > 20:
		fmt.Println("青年です")
	}
}
