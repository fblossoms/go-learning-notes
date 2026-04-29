package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 九九乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d × %d = %d\t\t", j, i, i*j)
		}
		fmt.Println()
	}

	a := "abcd测试"
	fmt.Println(len(a))
	// 当前长度为10（字符串长度/底层数组占用的总字节数），说明英文字符占1字节，汉字字符占3字节，Go能够识别出汉字，自动每三个字节读取
	// 该长度由Go使用utf-8（兼容ASCII）编译而来。
	// 底层的rune仍为int32（4字节）每英文字符，占不满置零对齐四字节
	for i, v := range a {
		fmt.Println(i, v) // v为对应的十进制的ASCII值。
		// 高级for循环按照字符遍历（有几个字符遍历几次），底层按照utf-8编码存储，返回unicode编码的码点，存在转换问题，由Go编译器负责
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	b := [3]int{1, 22, 333}
	for i, v := range b {
		fmt.Println(i, v)
	}
}
