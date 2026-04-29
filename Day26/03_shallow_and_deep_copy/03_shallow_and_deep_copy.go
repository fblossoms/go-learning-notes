package main

import (
	"fmt"

	"github.com/mohae/deepcopy" // 提供深拷贝的第三方库
)

func testCopy(s []int) {
	fmt.Printf("s1 %+v, %p, %p, %d, %d\n", s, &s, &s[0], len(s), cap(s))
}

func main() {
	s1 := []int{1, 3, 5}
	fmt.Printf("s1 %+v, %p, %p, %d, %d\n", s1, &s1, &s1[0], len(s1), cap(s1))
	s2 := s1 // 把header本质是大整数，所以是值传递
	fmt.Printf("s2 %+v, %p, %p, %d, %d\n", s2, &s2, &s2[0], len(s2), cap(s2))
	// 浅拷贝：只对地址进行复制，不对地址所指向的数据结构进行复制
	// 深拷贝：通过指针对指针所指向的数据结构（深入），复制时会创建数据结构的副本

	testCopy((s1)) // 依旧浅拷贝：高级语言拷贝均是浅拷贝

	// 深拷贝，Go官方不支持，需要通过第三方库使用
	x := deepcopy.Copy(s1)

	if v, ok := x.([]int); ok {
		v[0] = 100
		fmt.Println("s1:", s1)
		fmt.Println("v :", v)

		// 深拷贝，根据指针找到指针指向的最终数据类型进行拷贝，然后指针指向新的数据类型
		fmt.Printf("%T %[1]v\n", v)
		fmt.Printf("v %+v, %p, %p, %d, %d\n", v, &v, &v[0], len(v), cap(v))
	} else {
		fmt.Println("失败")
	}
}
