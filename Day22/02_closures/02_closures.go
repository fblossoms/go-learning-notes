package main

import "fmt"

func outer() func() {
	c := 65
	fmt.Println("0 c =", c, &c)
	inner := func() {
		c = 97 // 使用了外层函数的自由变量，所以是闭包
		fmt.Println("1 c =", c, &c)
	}
	return inner
}
func main() {
	f := outer() // f引用了outer()，所以保存了inner的地址
	f()
}
