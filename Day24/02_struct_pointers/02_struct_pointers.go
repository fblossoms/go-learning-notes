package main

import "fmt"

type Point struct {
	x, y int
}

func test(p Point) Point {
	fmt.Printf("%T, %[1]v\n, %p", p, &p)
	p.x = 2000
	fmt.Printf("%T, %[1]v\n, %p", p, &p)
	return p
}

func main() {
	p1 := Point{x: 10, y: 20} // Point实例
	fmt.Printf("%T, %[1]v\n", p1)
	p2 := &p1 // 指针
	fmt.Printf("%T, %[1]v\n", p2)
	p3 := new(Point) // new(类型) 为内建函数1
	fmt.Printf("%T, %[1]v\n", p3)

	// 访问属性
	fmt.Println(p1.x, p1.y)
	fmt.Println((*p2).x, (*p2).y)
	fmt.Println(p2.x, p2.y) // 相当于c++的p2 -> x语法

	// 成员写入
	p1.x = 100
	fmt.Println(p1.x, p1.y, p1)
	p2.x = 200 // 指向同一个实例，改的是同一个。改语法可以让程序员不用纠结是实例还是指针
	fmt.Println(p2.x, p2.y, p2)
	fmt.Println(p1.x, p1.y, p1)
	fmt.Println("=======================================")

	// 副本
	p5 := Point{x: 1, y: 4}
	fmt.Printf("%T, %[1]v %p\n", p5, &p5)

	p4 := p3
	fmt.Printf("%T, %[1]v %p\n", p4, &p4) // 并未使用同一个地址，因为短类型是值复制

	p6 := &p5
	fmt.Printf("%T, %[1]v %p\n", p6, &p6)

	p7 := p6
	fmt.Printf("%T, %[1]v %p\n", p7, &p7)
	fmt.Println("=======================================")

	pf1 := test(p5) // 如果想操作同一个数就传入&p5，此时函数的入参和出参也要改为指针类型
	// 看需求而定
	// 如果想操作不同，就使用实例，产生副本
	// 如果想操作相同，就是使用指针
	fmt.Printf("%T, %[1]v %p\n", p1, &p1)
	fmt.Printf("%T, %[1]v %p\n", pf1, &pf1)

	// pf2 := test(&p5)
	// fmt.Printf("%T, %[1]v %p\n", pf2, pf2)
	// pf2.y = 333
	// fmt.Printf("%T, %[1]v %p\n", p1, &pf1)
	// fmt.Printf("%T, %[1]v %p\n", pf1, pf1)
}
