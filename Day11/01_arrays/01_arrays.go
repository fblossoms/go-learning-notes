package main

import "fmt"

func showAddr(x [5]int) [5]int { // x是形式参数（形参）
	fmt.Printf("x %T, %[1]v, %p, %p, %p\n", x, &x, &x[0], &x[1])
	return x
}

func main() {
	var a1 [3]int // a1为标识符，类型为[3]int，零值定义顺序表，每一个元素类型为int（8byte），共占24字节
	var a2 [5]int // a2（[5]int）与a1（[3]int）类型不同（大类型都是数组，但具体类型不一样）
	fmt.Printf("%T, %T\n", a1, a2)

	fmt.Printf("%T %[1]v,\n %[3]p, %[4]p\n", a1, a1, &a1, &a1[0]) // 第一个元素的地址就是数组的地址

	// 初始化写法
	var a3 [3]int = [3]int{1, 2, 4}
	var a4 = [4]int{1, 3, 5, 7} // 推荐
	fmt.Println(a3, a4)

	// 不写默认零值
	var a5 = [5]int{1, 2}             // 可以只写一些值，后续再更改。但是不能后续新增元素个数
	var a6 = [5]int{0: 1, 4: 7, 2: 3} // 注意是索引值
	a5[0] = 99                        // 根据索引定位覆盖原值
	fmt.Println(a5, a6)

	// 一维遍历
	for i := 0; i < len(a6); i++ {
		fmt.Print(a6[i])
	}
	for i, v := range a6 {
		fmt.Printf("\n%d: %d", i, v)
	}

	var a7 = [...]int{11, 22, 33} // 不写长度使用就根据定义的元素个数推断
	var a8 = [...]int{1: 11, 22, 8: 33}
	fmt.Println(a7, a8)

	// n维（工程一般最多2维）
	var a9 = [3][4]int{{100, 200, 11, 22}, {100, 222}} // 逗号（,）分隔定义每行元素的值，同样可以使用根据索引初始化和[...]
	a9[0][1] = 11                                      // 多维根据行和列索引修改
	fmt.Println(a9, a9[2], a9[2][2], len(a9), cap(a9)) // a9[2]为行索引，对应第3行。a9[2][2]为行和列索引，对应第3行第3列
	fmt.Printf("\n%T", a9)                             // 类型为[3][4]int

	// 二维遍历（嵌套）
	for i := 0; i < len(a9); i++ {
		fmt.Printf("第 %d 行", i)
		for j := 0; j < len(a9[i]); j++ {
			fmt.Printf("%d  ", a9[i][j])
		}
		fmt.Println()
	}
	for i, row := range a9 {
		fmt.Printf("第 %d 行", i+1)
		for _, val := range row {
			fmt.Printf("%d  ", val)
		}
		fmt.Println()
	}

	// 字符串
	var b1 = [3]string{"abc", "123456789012345678901234567890"}
	fmt.Println(b1)
	fmt.Printf("%p\n%p\n%p\n", &b1[0], &b1[1], &b1[2])

	// 地址
	var c = [5]int{1, 2, 3, 4, 5}
	var d = c
	fmt.Println(c == d, &c == &d) // 内容一样，地址不一样（区别于Python）。相当于拷贝了一份，工程中需要考虑内存的问题
	// showAddr(c)                   // 使用函数传参地址也不一样，也是使用了值拷贝
	// e := showAddr(c)				 // 使用返回值仍会发生值拷贝
}
