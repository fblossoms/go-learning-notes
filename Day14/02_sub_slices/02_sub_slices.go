package main

import "fmt"

func main() {
	s1 := []int{10, 30, 50, 70, 90} // 初始化
	fmt.Printf("&s1=%p, &s1[0]=%p, len=%-2d, cap=%-2d, s1=%v\n", &s1, &s1[0], len(s1), cap(s1), s1)

	s2 := s1[:] // 不写start和end就相当于完全拷贝，此时共用底层数组
	fmt.Printf("&s2=%p, &s2[0]=%p, len=%-2d, cap=%-2d, s2=%v\n", &s2, &s2[0], len(s2), cap(s2), s2)

	s3 := s1[1:] // 不写end相当于默认从start索引值取到最后一个（掐头），此时不共用底层数组，len和cap都改变
	fmt.Printf("&s3=%p, &s3[0]=%p, len=%-2d, cap=%-2d, s3=%v\n", &s3, &s3[0], len(s3), cap(s3), s3)

	s4 := s1[1:4] // 前包后不包，所以取到索引3，此时len改变，cap不变与s3共用底层数组
	fmt.Printf("&s4=%p, &s4[0]=%p, len=%-2d, cap=%-2d, s4=%v\n", &s4, &s4[0], len(s4), cap(s4), s4)

	s5 := s1[2:4]
	fmt.Printf("&s5=%p, &s5[0]=%p, len=%-2d, cap=%-2d, s7=%v\n", &s5, &s5[0], len(s5), cap(s5), s5)

	s6 := s1[:4]
	fmt.Printf("&s6=%p, &s6[0]=%p, len=%-2d, cap=%-2d, s6=%v\n", &s6, &s6[0], len(s6), cap(s6), s6)

	s7 := s1[1:1]       // len=0，cap=4，ptr从索引2开始，且为空切片，空切片无法打印
	s7 = append(s7, -1) // 追加元素后可以打印
	fmt.Printf("&s7=%p, &s7[0]=%p, len=%-2d, cap=%-2d, s7=%v\n", &s7, &s7[0], len(s7), cap(s7), s7)

	//s8 := s1[4:4] // len=0，cap=1 // cap计算方法为从ptr指向的首地址指向最后一元素
	//fmt.Printf("&s8=%p, &s8[0]=%p, len=%-2d, cap=%-2d, s8=%v\n", &s8, &s8[0], len(s8), cap(s8), s8)

	s9 := s1[5:5] // cap为0时，ptr指向底层数组首地址，为了安全
	//fmt.Printf("&s9=%p, &s9[0]=%p, len=%-2d, cap=%-2d, s9=%v\n", &s9, &s9[0], len(s9), cap(s9), s9)
	print(s9)
}
