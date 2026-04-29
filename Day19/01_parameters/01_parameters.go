package main

import (
	"fmt"
)

func fn1()                   {} // 无形参
func fn2(int)                {} // 有一个int形参，但是没办法用到，不推荐
func fn3(x int)              {} // 单参函数。形式参数是局部变量
func fn4(x int, y int)       {} // 多参函数
func fn5(x, y int, z string) {} // 相邻参数类型相同，可以写到一起

// 可变形参
func t1(nums ...int) { // ...相当于等等等，后续可以传入0个或多个实参
	fmt.Printf("%T\n", nums)                               // 类型为切片
	fmt.Printf("%v, %d, %d\n", nums, len(nums), cap(nums)) // 提供的一个个实参收集起来放在nums切片中
}

// func t2(x, y int, nums ...int, z ...int) // 混合使用时，可变形参只能放在形参列表的最后一位
func t2(x, y int, nums ...int) {
	fmt.Println(x, y, "|||", nums)
	fmt.Printf("%p, %p, %d, %d\n\n", &nums, &nums[0], len(nums), cap(nums))
}

// 对比
func f1(nums []int)  {} // 只能传切片
func f2(nums ...int) {} // 可以传参可以是切片，也可以是要组成的切片

func main() {
	fn1()
	fn2(10)
	fn3(20)
	fn4(40, 60)
	fn5(10, 20, "abc") // 有几个形参就写多少个实参，不能缺省，没有默认值

	// 可变参数
	t1(100)
	a1 := []int{1, 2, 3, 4}
	// t6(a1) // 错误。虽然形参...int类型是切片，但是不支持这种写法
	t1(a1...)          // 正确。定义时使用...Type时，传参时就要写 切片名...
	t1([]int{1, 2}...) // Go支持在可变参数上使用该语法
	// fn4([]int{1, 2}...)	// Go不支持在普通参数上使用该语法

	// t2()									// 错误，不支持缺省值
	// t2(1)								// 错误，不支持缺省值
	// t2(1, 2)								// 正确，因为切片长度可以为0
	// t2(1, 2, 3, 4, []int{5, 6, 7}...)	// 错误，不支持混合传参

	fmt.Println("-----------------------------------------------------")
	t2(1, 2, 3) // 正确，会自动构建切片
	s1 := []int{3}
	t2(1, 2, s1...) // 当传入切片时，不拆开切片传入实参，直接共用同一个底层数组。不同的slice header，共用一个底层数组
	fmt.Printf("%p, %p, %d, %d\n\n", &s1, &s1[0], len(s1), cap(s1))
	fmt.Println("-----------------------------------------------------")
	t2(1, 2, 3, 4) // 正确，剩下的识别为切片
	s2 := []int{3, 4}
	t2(1, 2, s2...) // 不同的slice header，共用一个底层数组
	fmt.Printf("%p, %p, %d, %d\n\n", &s2, &s2[0], len(s2), cap(s2))
	fmt.Println("-----------------------------------------------------")
	t2(1, 2, []int{3, 4, 5}...) // 正确

	// ...
	s := []any{1, 2, 3}
	// for i, v := range a{}
	fmt.Println(s)    // 一个元素按照位置给a提供[]any{t}
	fmt.Println(s...) // ...共用底层数组，a和t共用底层数组
}
