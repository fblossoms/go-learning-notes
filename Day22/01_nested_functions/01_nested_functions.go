package main

import "fmt"

func outer1() {
	c := 65
	inner := func() {
		fmt.Println("1 c =", c)
	}
	// inner（函数型）和 c（整型）都是outer1的局部变量，只是类型不一样
	inner() // 调用，依旧有压栈操作，压在outer1的栈帧上
	fmt.Println("2 c =", c)
}

func outer2() {
	c := 65
	inner := func() {
		c = 97
		fmt.Println("1 c =", c)
	}
	// inner（函数型）和 c（整型）都是outer1的局部变量，只是类型不一样
	inner() // 调用，依旧有压栈操作，压在outer1的栈帧上
	fmt.Println("2 c =", c)
}

func outer3() {
	c := 65
	fmt.Println("0 c =", c, &c)
	inner := func() {
		c = 97
		fmt.Println("1 c =", c, &c)
		c := c + 1
		fmt.Println("3 c =", c, &c)
	}
	fmt.Printf("%T, %p\n", inner, &inner)
	inner()
	fmt.Println("2 c =", c, &c)
}

func main() {
	outer1()
	outer2()
	// inner() // 错误，因为这是outer1的局部变量，对外不可见
	for i := 0; i < 5; i++ { // 每次调用不会共享内存
		outer3()
	}
}
