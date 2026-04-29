package main

import "fmt"

type Person struct {
	name string
	age  int
}

// 隐式实现的接口方法
// 只读取，返回string类型，对#v没有影响
//func (Person) String() string {
//	return "abc"
//}

// 只读取，返回string类型，对#v没有影响。对非指针没有影响
//func (*Person) String() string {
//	return "abcd"
//}

// GoString()影响#v，实例和指针
//func (Person) GoString() string {
//	return "abcd"
//}

// 影响#v，指针
func (*Person) GoString() string {
	return "abcd"
}

// 如果对int打印不满意————定义新类型
// 用于想让一个类型具有方法，就定义新类型扩展方法
type MyInt int

func (i MyInt) String() string {
	return fmt.Sprintf("我是一个int打印的扩展 = %d", i)
}

func main() {
	m := Person{"Tom", 20}
	fmt.Println(1, m, &m) // Println默认用%v打印
	fmt.Printf("2 %v, %v\n", m, &m)
	fmt.Printf("3 %+v, %+v\n", m, &m)
	fmt.Printf("4 %#v, %#v\n", m, &m)

	// var i = 123	// 无效打印
	var i MyInt = 123
	fmt.Println(i)
}
