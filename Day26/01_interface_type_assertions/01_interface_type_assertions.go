package main

import "fmt"

func main() {
	var x any
	n := "abc"
	x = n
	fmt.Printf("%T %[1]v\n", x)

	// t := x.(int) // 断言失败就报错
	t := x.(string) // 断言成功，则可使用t，此时t为string类型，可以进行有关string的操作
	fmt.Printf("%T %[1]v\n", t)
	fmt.Println(t + "xyz") // 可以进行有关string的操作，若未进行断言则仍为any类，此时无法进行字符串相加操作

	//if m, ok := x.(string); ok {
	//	fmt.Printf("断言成功：%T %[1]v, %T %[2]v\n", m, ok)
	//} else {
	//	fmt.Printf("断言失败：%T %[1]v, %T %[2]v\n", m, ok)
	//}

	// switch写法，用于做多种接口的断言。Go语言的case不会穿透
	switch x.(type) {
	case nil:
		fmt.Printf("%T %[1]v", x)
	case int:
		fmt.Printf("%T %[1]v", x)
	case []int:
		fmt.Printf("%T %[1]v", x)
	default:
		fmt.Printf("%T %[1]v", x)
	}

	switch y := x.(type) {
	case nil:
		fmt.Printf("%T %[1]v", y)
	case int:
		fmt.Printf("%T %[1]v", y)
	case []int:
		fmt.Printf("%T %[1]v", y)
	default:
		fmt.Printf("%T %[1]v", y)
	}
}
