package main

import "fmt"

/*
顶层代码区：可定义全局标识符（const全局常量、var全局变量，短格式写法不可用）
*/
func main() {
	// const a = "gxedu"        // 用到类型推导
	const a string = "gxedu" // 没有用到类型推导
	fmt.Println(a)

	const b string = "100" // 左边指定了（字符串）类型就要严格按照类型要求写右边，此时右边应改为"100"
	fmt.Println(b)

	// const c int 	// 常量必须初始化，必须在当行给出具体初始值
	var c int // 变量可以延迟赋值，没有初始化时会自动赋零值（例如int为0，float64为0等）
	fmt.Println(c)

	// 常量批量初始化：
	// const d int, e int = 1, 2	// 不可以
	const d, e int = 1, 2 // 可以
	fmt.Println(d, e)
	// const f string, g int = "1", 2	// 不可以
	const ( // 可以
		d2 string = "1"
		e2 int    = 2
	)
	fmt.Println(d2, e2)

	// 变量批量初始化：
	// var f int, g int = 1, 2		// 不可以
	var f, g int = 1, 2 // 可以
	fmt.Println(f, g)
	//var f2 string, g2 int = "1", 2	// 不可以
	var ( // 可以
		f2 string = "1"
		g2 int    = 2
	)
	fmt.Println(f2, g2)

	// 零值
	var (
		h int
		i float64
		j string
		k bool
	)
	fmt.Println(h, i, j, k)

	// 数组
	var l = [3]int{11, 22, 33} // [3]int，3表示声明3个元素，数组元素的类型按照定义严格相同（这里是int），{}里写元素，
	// Go中数组长度不可变，元素能变
	// const l2 = [3]int{11, 22, 33} 	// 常量不可以定义数组类型，因为Go认为数组长度会变不符合常量定义
	l3 := [3]int{44, 55, 66}
	fmt.Println(l, l3)

	// 短格式
	m, n := 5, "34"
	fmt.Println(m, n)

	// _ 下划线空白标识符：此时"ok"进入_直接作废
	p, _, q := func() (int, string, bool) {
		return 300, "ok", true
	}()
	fmt.Println(p, q)

	// 两数交换：仅限于同类型
	// 写法1：引入中间值
	var r, r2 int = 11, 11
	var s, s2 int = 22, 22
	t := r
	s = t
	r = s
	fmt.Println(r, s)
	// 写法2：直接换
	r2, s2 = s2, r2
	fmt.Println(r2, s2)
}
