package main

import "fmt"

type Animal struct {
	name string
	age  int
}

// 用该函数构建一个全新的实例，即构造函数（普通函数，不是该类型的方法）
// 构造函数往往返回该类型的指针
func NewAnimal() Animal { // 习惯会命名New前缀（想要包外可见时）和new前缀（想要包外不可见时）
	t := Animal{}
	fmt.Printf("%T, %[1]v, %p\n", t, &t)
	return t
}

func NewAnimal2() *Animal { // 习惯会命名New前缀（想要包外可见时）和new前缀（想要包外不可见时）
	t := Animal{}
	fmt.Printf("%T, %[1]v, %p\n", t, &t)
	return &t
}

func NewAnimal3() *Animal {
	t := new(Animal) // new(type)用来分配内存并返回指针
	fmt.Printf("%T, %[1]v, %p\n", t, t)
	return t
}

// Go不支持函数重载overload
//func NewAnimal(name string, age int) *Animal {
//	// return &Animal{name, age} // 形参，注意要按顺序写
//	return &Animal{name: name, age: age} // 属性名: 形参
//}

// 构造

func main() {
	// var a1 Animal            // 定义了一个实例，但是由于没有赋值，所以为零值
	// var a2 = Animal{age: 30} // 未赋值的默认零值
	var a3 = NewAnimal()
	var a4 = NewAnimal2()
	var a5 = NewAnimal3()

	fmt.Printf("%T, %[1]v, %p\n", a3, &a3) // 地址不一致，产生了值复制
	fmt.Printf("%T, %[1]v, %p\n", a4, a4)  // 地址一致，未产生值复制
	fmt.Printf("%T, %[1]v, %p\n", a5, a5)  // 地址一致，未产生值复制
}
