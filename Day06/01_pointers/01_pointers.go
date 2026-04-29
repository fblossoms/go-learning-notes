package main

import "fmt"

func main() {
	fmt.Println(5 % 2) // 取模 mod
	fmt.Println(100 % 5)

	fmt.Println(2<<1, 2*2, 1<<3, 1*8, 3<<1, 3*2) // 左移几位相当于乘2的几次方
	fmt.Println(2>>1, 2/2, 1>>3, 1/8, 3>>1, 3/2) // 右移几位相当于除2的几次方

	fmt.Println(2 | 3) // 不进位

	fmt.Println(15&5, 15&^5)

	fmt.Println(1 > 1.2 || "abc" != "abc" || 5 > 3)

	var a = 100
	a += 100 + 200 // 等价于a += (100 + 200)即 a = a + 300
	fmt.Println(a)
}
