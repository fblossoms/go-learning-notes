package main

import "fmt"

func main() {
	s0 := []int{10, 30, 50, 70, 90, 110, 130, 150}
	fmt.Printf("&s0=%p, &s0[0]=%p, len=%-2d, cap=%-2d, s0=%v\n", &s0, &s0[0], len(s0), cap(s0), s0)

	s1 := make([]int, 3)
	fmt.Printf("&s1=%p, &s1[0]=%p, len=%-2d, cap=%-2d, s1=%v\n", &s1, &s1[0], len(s1), cap(s1), s1)

	n := copy(s1, s0) // 从s0里元素，按索引拷贝拷贝给到s1
	fmt.Printf("&s0=%p, &s0[0]=%p, len=%-2d, cap=%-2d, s0=%v\n", &s0, &s0[0], len(s0), cap(s0), s0)
	fmt.Printf("&s1=%p, &s1[0]=%p, len=%-2d, cap=%-2d, s1=%v\n", &s1, &s1[0], len(s1), cap(s1), s1)

	//s2 := make([]int, 0) // 初始化时len=cap=0时，则无法进行拷贝，有底层数组
	//var s2 []int // 初始化时len=cap=0时，则无法进行拷贝，因为指针ptr指向nil
	s2 := make([]int, 1) // 有长度时才能进行拷贝
	fmt.Printf("&s2=%p, len=%-2d, cap=%-2d, s2=%v\n", &s2, len(s2), cap(s2), s2)
	//copy(s2, s0)
	copy(s2, s0[2:4]) // 参数为切片，不能单独使用索引（如s0[2]）取特定值进行拷贝，可以使用子切片进行指定（如s0[2:3]）
	fmt.Printf("&s2=%p, len=%-2d, cap=%-2d, s2=%v\n", &s2, len(s2), cap(s2), s2)

	fmt.Println(n)

	// 合并1，不能在现有切片上合并，底层数组不能扩大，且影响原数据
	s3 := make([]int, len(s0)+len(s1)) // 只关心有效数据长度，所以取的是len，不取cap
	copy(s3, s0)
	fmt.Printf("&s3=%p, &s3[0]=%p, len=%-2d, cap=%-2d, s3=%v\n", &s3, &s3[0], len(s3), cap(s3), s3)
	// copy(s3, s2) // 未起到合并效果，仅替换元素
	//fmt.Printf("&s3=%p, &s3[0]=%p, len=%-2d, cap=%-2d, s3=%v\n", &s3, &s3[0], len(s3), cap(s3), s3)
	copy(s3[len(s0):], s1) // 此时可以起到合并效果
	fmt.Printf("&s3=%p, &s3[0]=%p, len=%-2d, cap=%-2d, s3=%v\n", &s3, &s3[0], len(s3), cap(s3), s3)

	// 合并2
	s4 := make([]int, 0, len(s0)+len(s1))
	// 注意不要写成make([]int, len(s0)+len(s1))，此时尾部追加会扩容。也不要写成make([]int, 0)，此时容量没有一次给够
	fmt.Printf("&s4=%p, len=%-2d, cap=%-2d, s3=%v\n", &s4, len(s4), cap(s4), s4)
	s4 = append(s4, s0...) // append的参数不仅可以是元素值，也可以是切片，注意追加的切片后要写...
	fmt.Printf("&s4=%p, &s4[0]=%p, len=%-2d, cap=%-2d, s3=%v\n", &s4, &s4[0], len(s4), cap(s4), s4)
	s4 = append(s4, s1...) // append的参数不仅可以是元素值，也可以是切片，注意追加的切片后要写...
	fmt.Printf("&s4=%p, &s4[0]=%p, len=%-2d, cap=%-2d, s3=%v\n", &s4, &s4[0], len(s4), cap(s4), s4)

	// delete（若要删除切片中间段区间的子切片，实际思路转换为新开辟一段空间，放上不需要删除的子切片。效率不高）
	fmt.Printf("&s0=%p, &s0[0]=%p, len=%-2d, cap=%-2d, s0=%v\n", &s0, &s0[0], len(s0), cap(s0), s0)
	// var s5 []int	// 此写法后续需要扩容次数过多
	s5 := make([]int, 0, 4) // 根据自己目标需要的个数来开辟空间
	s5 = append(s5, s0[:2]...)
	fmt.Printf("&s5=%p, &s5[0]=%p, len=%-2d, cap=%-2d, s5=%v\n", &s5, &s5[0], len(s5), cap(s5), s5)
	s5 = append(s5, s0[6:]...)
	fmt.Printf("&s5=%p, &s5[0]=%p, len=%-2d, cap=%-2d, s5=%v\n", &s5, &s5[0], len(s5), cap(s5), s5)
}
