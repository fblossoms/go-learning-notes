package main

import (
	"fmt"
)

// 类型：认为该类型都必须具备的特征
type Animal struct {
	// 定义属性
	name string
	age  int
	// Age int	// 包外可见也是一种访问控制
}

type Runner interface {
	run()
}

// 定义方法
func (a *Animal) run() {
	fmt.Println("动物 跑")
	fmt.Printf("动物 跑，type=%T，v=%[1]v\n", a) // 类型为*Animal
	// 等价于a *Animal = &c（复习，也可以a Animal = c，只是多了份拷贝）
	// 两者类型不匹配，但是可以赋值，因为Cat是Animal的子类（c属于Cat类，也始于Animal类）
	// t Cat = Animal{}	// 不可以，某动物属于Animal类，就属于Cat类，即不可以像这样反着说
}

// 继承：实现Cat和Animal的关系
type Cat struct {
	Animal        // 嵌套进去的Animal就算是Cat的父类parent
	color  string // 自我特征
}

func (*Cat) leap() {
	fmt.Println("猫 跳")
}

// 覆盖：不满意父类的run方法，自己实现
func (c *Cat) run() {
	// 锦（父类）上添花（子类），首先要使用父方法打个底
	// c.run() // 无限递归了
	c.Animal.run()

	fmt.Println("猫 跑")
	fmt.Printf("猫 跑，type=%T，v=%[1]v\n", c) // 类型为*Cat
	// 自己有就用自己的
}

func main() {
	a := Animal{name: "Tibbers", age: 3} // 实例化
	fmt.Println(a.name, a.age)           // 访问实例的属性
	a.run()
	// a.leap()	// a不能调cat的自有方法

	c := Cat{Animal: Animal{name: "Yuumi", age: 3}, color: "white"} // 继承
	fmt.Println(c.name, c.Animal.age, c.color, c)                   // 访问实例的属性
	c.run()                                                         // 用自己的
	// c.Animal.run()                                                  // 此时用的是父类

	// var a Animal = c	// Java可以。Go中不能这样赋值，Go中需要使用接口来间接实现
	var t Runner = &c // 改写法不要求有父子关系，有子类就用子类，子类没有就用子类的父类。如果是指针接收就用&标识符（如果是*Cat就是&c）
	t.run()
	// t = d // 假设d为另一实例
	// t.Animal.run	// 错误t是接口类型变量，不知道什么是Animal

}
