package main

import "fmt"

// 如何使用匿名结构体
type Point1 struct { // type 定义新类型
	x, y int
}

func test1(x struct{ x, y int }) { // 一般不这样写
	fmt.Printf("%T, %[1]v\n", x)
}
func test2(x Point1) { // 一般这样写
	fmt.Printf("%T, %[1]v\n", x)
}

// 匿名成员，不经常使用，因为名称要有含义
type Point2 struct {
	x   int
	int // 此时，成员名和类型均为int，所以同类型只能写一次匿名成员，注意赋值时也必须按照类型要求
	bool
}

func main() {

	// 如何使用匿名结构体
	var t1 struct { // var 定义变量（本质是实例），该变量有类型（类型是struct和{}里所有语句，有变量名，这里是point），但是没有定义类型，类型匿名
		x, y int
	} // 使用变量无法重复使用，因为没有名字
	var t2 struct{ x, y int }
	var t3 = struct {
		x, y int
	}{x: 30, y: 40} // 相当于Point{}，只是这里是匿名（只能一次性使用）
	var p1 Point1 // 先定义Point类型，再实例化给了p1

	fmt.Printf("%T, %[1]v\n", t1)
	fmt.Printf("%T, %[1]v\n", t2)
	fmt.Printf("%T, %[1]v\n", p1)
	fmt.Printf("%T, %[1]v\n", t3)

	test1(t1)
	test1(t2)
	test1(p1)

	// 匿名成员
	p2 := Point2{}
	fmt.Printf("%T, %#[1]v\n", p2)
	fmt.Println(p2.x, p2.int)
}
