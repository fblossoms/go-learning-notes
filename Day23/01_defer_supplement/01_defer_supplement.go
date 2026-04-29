package main

import "fmt"

var count = -1

func test1() int {
	count = 1
	defer func() {
		count++
		fmt.Println(3, count)
	}()
	return count // 值传递
}

func test2() (count int) {
	count = 1
	defer func() {
		count++
		fmt.Println(3, count)
	}()
	return count // 值传递
}

func test3() int {
	count := 100
	defer func() {
		count++
		fmt.Println(3, count)
	}()
	return count // 值传递
}

func main() {
	fmt.Println(1, test1())
	fmt.Println(2, count)
	fmt.Println("-----------------------------------")
	fmt.Println(1, test2())
	fmt.Println(2, count)
	fmt.Println("-----------------------------------")
	fmt.Println(1, test3())
	fmt.Println(2, count)
}
