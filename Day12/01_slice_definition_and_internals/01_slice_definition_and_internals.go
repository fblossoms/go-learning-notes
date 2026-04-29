package main

import "fmt"

func main() {
	var a = [3]int{1, 2, 3} // 定义数组，注意[...]int{1, 2, 3}也是数组
	var b = []int{1, 2, 3}  // 定义切片，类型为[]int，{}表示字面量开始了
	fmt.Printf("len=%-2d\ncap=%-2d\n", len(b), cap(b))
	fmt.Println(a, b)

	// 可以用索引指定
	var c = []int{1: 20, 2, 3} // 定义切片，类型为[]int，{}表示字面量开始了
	fmt.Println(c)

	// 若未赋初值
	var d []int           // 没有默认值，就给零值nil，给的是标头值
	fmt.Println(d == nil) // 标头值有了，但ptr悬空（nil）说明没有底层数组，len=0，cap=0
	fmt.Printf("len=%-2d\ncap=%-2d\n", len(d), cap(d))

	var e = []int{}
	fmt.Println(e == nil) // 初始化了，有标头值，分配了底层数组，ptr指向了数组
	fmt.Printf("len=%-2d\ncap=%-2d\n", len(e), cap(e))

	s1 := make([]int, 10) // make（内建函数）可以定义slice、map、channel三种类型。第一个参数写类型。此时可以用索引s1[0]。s1[10]超界
	// s2 := make([]int, 0)  // 相当于var s2 = []int{}，切片len=cap=0，底层数组len=cap=0。受len影响，此时不可以用索引s2[0]
	// s3 := make([]int, 0, 10)  // ptr有底层数组，切片header中len=0，cap=1，底层数组len=cap=10。受len影响，此时不可以用索引s3[0]
	fmt.Printf("len=%-2d\ncap=%-2d\n", len(s1), cap(s1))
	
}
