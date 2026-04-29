package main

import "fmt"

// 1、传统方法
func fibLoop(n int) int {
	switch {
	case n < 0:
		panic("n is negative")
	case n == 0:
		return 0
	case n == 1 || n == 2:
		return 1
	}
	a, b := 1, 1
	for i := 0; i < n-2; i++ {
		a, b = b, a+b
	}
	return b
}

// 递归
func fib1(n int) int {
	if n < 0 {
		panic("n is negative")
	} else if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)

}

// 递归方法2
func fib2(n, a, b int) int {
	if n < 0 {
		panic("n is negative")
	} else if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return b
	}
	a, b = b, a+b
	return fib2(n-1, a, b)
}

func main() {
	fmt.Println(fibLoop(50))
	fmt.Println(fib1(50))
	fmt.Println(fib2(50, 1, 1))

}
