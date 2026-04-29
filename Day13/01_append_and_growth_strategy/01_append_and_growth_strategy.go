package main

import "fmt"

func main() {
	// append
	var a []int
	fmt.Printf("%p, l=%d, c=%d, %v\n", &a, len(a), cap(a), a) // ptr -> nil
	a = append(a, 10)
	fmt.Printf("%p, %p, l=%d, c=%d, %v\n", &a, &a[0], len(a), cap(a), a) // ptr -> array

	// 赋值
	s1 := []int{10}
	fmt.Printf("%p, l=%d, c=%d, %v\n", &s1, len(s1), cap(s1), s1)
	s1 = make([]int, 2, 5) // 地址不变，prt指向新的切片
	fmt.Printf("%p, l=%d, c=%d, %v\n", &s1, len(s1), cap(s1), s1)
	s1 = append(s1, 100) // &s1是header的地址，&s1[0]是底层数组的地址
	fmt.Printf("%p, %p, l=%d, c=%d, %v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	fmt.Println("----------------------------------------------------------------------")
	// s1 = append(s1, 200, 300) // header更新时会覆盖原header
	// fmt.Printf("%p, %p, l=%d, c=%d, %v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	s2 := append(s1, 1, 2) // s1不变（因为header不同），s2有自己的header，但header的ptr指向s1的底层数组（s1和s2共用一个底层数组）
	fmt.Printf("%p, %p, l=%d, c=%d, %v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	fmt.Printf("%p, %p, l=%d, c=%d, %v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	fmt.Println("----------------------------------------------------------------------")
	s3 := append(s1, -1) // 由于s1、s2、s3共用一个底层数组，所以新append时会覆盖原append值
	fmt.Printf("%p, %p, l=%d, c=%d, %v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	fmt.Printf("%p, %p, l=%d, c=%d, %v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	fmt.Printf("%p, %p, l=%d, c=%d, %v\n", &s3, &s3[0], len(s3), cap(s3), s3)
	fmt.Println("----------------------------------------------------------------------")
	s4 := append(s3, 3, 4, 5) // s4追加后超出s3原长度，自动开辟新底层数组，因为数组长度不可变
	fmt.Printf("%p, %p, l=%d, c=%d, %v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	fmt.Printf("%p, %p, l=%d, c=%d, %v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	fmt.Printf("%p, %p, l=%d, c=%d, %v\n", &s3, &s3[0], len(s3), cap(s3), s3)
	fmt.Printf("%p, %p, l=%d, c=%d, %v\n", &s4, &s4[0], len(s4), cap(s4), s4)
	fmt.Println("----------------------------------------------------------------------")
	s5 := append(s4, 6, 7, 8, 9, 10, 11) // len>cap，开辟新底层数组，容量cap * 2
	fmt.Printf("%p, %p, l=%d, c=%d, %v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	fmt.Printf("%p, %p, l=%d, c=%d, %v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	fmt.Printf("%p, %p, l=%d, c=%d, %v\n", &s3, &s3[0], len(s3), cap(s3), s3)
	fmt.Printf("%p, %p, l=%d, c=%d, %v\n", &s4, &s4[0], len(s4), cap(s4), s4)
	fmt.Printf("%p, %p, l=%d, c=%d, %v\n", &s5, &s5[0], len(s5), cap(s5), s5)
	fmt.Println("----------------------------------------------------------------------")

}
