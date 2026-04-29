package main

import (
	"fmt"
)

func showAddr(x []int) {
	fmt.Printf("x %p, %p, len=%-2d, cap=%-2d, %v\n", &x, &x[0], len(x), cap(x), x) // x也会拷贝一份slice header
	if len(x) > 0 {                                                                // 影响共同使用的底层数组
		x[0] = 100
	}
	fmt.Printf("x %p, %p, len=%-2d, cap=%-2d, %v\n", &x, &x[0], len(x), cap(x), x) // x也会拷贝一份slice header

}

func main() {
	s1 := []int{1, 2, 3} // 初始化
	fmt.Printf("s1 %p, %p, len=%-2d, cap=%-2d, %v\n", &s1, &s1[0], len(s1), cap(s1), s1)

	s2 := s1 // 赋值（引用）
	fmt.Printf("s2 %p, %p, len=%-2d, cap=%-2d, %v\n", &s2, &s2[0], len(s2), cap(s2), s2)

	showAddr(s1)
	fmt.Printf("s1 %p, %p, len=%-2d, cap=%-2d, %v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	fmt.Printf("s2 %p, %p, len=%-2d, cap=%-2d, %v\n", &s2, &s2[0], len(s2), cap(s2), s2)
}
