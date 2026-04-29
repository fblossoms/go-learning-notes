package main

import "fmt"

type PointR struct {
	x, y int
}

// 通过方法取成员
// 一个属性只提供了getter方法，只读属性
func (p PointR) GetX() int {
	fmt.Printf("%+v, %p", p, &p)
	return p.x
}

func (p *PointR) GetY() int {
	fmt.Printf("%+v, %p", p, p)
	return p.y
}

// setter 传参，无返回值
func (p PointR) SetX(value int) {
	p.x = value // 传入值
}

// Set必须用指针
func (p *PointR) SetY(value int) {
	p.x = value // 传入值
}

// 无receiver名
// 如果在一个方法中，想强调foo()不是一个普通的函数，只是该PointR类型的方法，又不需要使用实例或指针，内部并不需要通过recevier来使用它的属性，不需要访问它的成员
// 此时就能省略为一个类型或一个类型的指针
func (PointR) foo() {

}

func main() {
	p1 := PointR{4, 5}
	fmt.Printf("p1 %T %[1]v %p %d %d\n", p1, &p1, p1.x, p1.y)
	//p1.x = 100

	//p2 := &p1
	//fmt.Println(p1.x, p1.GetX())
	//fmt.Println(p2.x, p2.GetX()) // 通过指针访问（*Point），发生值复制
	//fmt.Println("=================================================")
	//
	//fmt.Println(p1.y, p2.y)
	//fmt.Println(p1.GetY()) // 直接访问地址里的值，取p1的地址
	//fmt.Println(p2.GetY()) // 如果实例是类型的地址，就把地址直接传过去

	p1.SetX(200) // 改的是这份副本的p.x，原来的p1.x 不变
	fmt.Printf("p1 %T %[1]v %p %d %d\n", p1, &p1, p1.x, p1.y)

	p2 := &p1
	p2.SetX(300)
	fmt.Printf("p1 %T %[1]v %p %d %d\n", p1, &p1, p1.x, p1.y)

	p1.SetY(200) // 改的是同一个地址里面放的值p.x，原来的p1.x改变
	fmt.Printf("p1 %T %[1]v %p %d %d\n", p1, &p1, p1.x, p1.y)

	//p2 := &p1
	//p2.SetY(300)
	//fmt.Printf("p1 %T %[1]v %p %d %d\n", p1, &p1, p1.x, p1.y)
}
