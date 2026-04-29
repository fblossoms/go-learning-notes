package main

import "fmt"

func main() {
	// defer fmt.Println("1")
	// fmt.Println("main start")
	// defer fmt.Println("2")
	// defer fmt.Println("3") // 先注册的后执行，后注册的先执行
	// fmt.Println("main end")
	// return

	// defer fmt.Println("1")
	// fmt.Println("main start")
	// defer fmt.Println("2")
	// panic("我出错了") // 结束，下面的语句不会执行
	// defer fmt.Println("3")
	// fmt.Println("main end")
	// return

	// count := 1
	// defer fmt.Println(count) // 注册时计算，保留当前的状态
	// count++
	// fmt.Println("s")
	// defer fmt.Println(count)
	// count++
	// defer fmt.Println(count)
	// fmt.Println("e")

	count := 1
	defer func() {
		fmt.Println(count)
	}()
	count++
	fmt.Println("s")
	defer func() {
		fmt.Println(count)
	}()
	count++
	defer func() {
		fmt.Println(count) // 没有count传入，默认用外部函数的count值，所以均为3
	}()
	fmt.Println("e")

}
