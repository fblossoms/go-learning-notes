package main

import "fmt"

// 臃肿写法，需要定义重复变量
//type Animals struct {
//	name string
//	age  int
//}
//type Cat struct {
//	name  string
//	age   int
//	color string
//}
//type Dog struct {
//	name string
//	age  int
//	sex  string
//}

// 结构体嵌套写法，较为麻烦
//type Animals struct {
//	name string
//	age  int
//}
//type Cat struct {
//	a     Animals
//	color string
//}

// 父子结构（继承）写法，本质是使用匿名函数
type Animals struct { // 父类
	name string
	age  int
}
type Cat struct { // 子类
	Animals
	age   int // 虽然父类已提供，但是优先用自己定义的
	color string
}

func main() {
	// 臃肿写法
	// var c1 = Cat{"xiaobai", 3, "black"}
	// fmt.Printf("%+v\n", c1)

	// 结构体嵌套写法
	// var c2 = Cat{Animals{"xiaobai", 3}, "black"} // 结构体嵌套写法赋值
	// fmt.Printf("%+v\n", c2)
	// fmt.Println(c2.a.name, c2.a.age, c2.color)
	// c2.a.name = "xiaohei" // 改变
	// c2.a.age = 4
	// c2.color = "white"
	// fmt.Println(c2.a.name, c2.a.age, c2.color)

	// 父子结构写法
	var c3 = Cat{}
	fmt.Println(c3.Animals.name, c3.Animals.age, c3.color)
	fmt.Println(c3.name, c3.age, c3.color) // 两种为相同写法。注意严格来说，c3里面并没有name和age（是Animal的），相当于是语法糖
}
