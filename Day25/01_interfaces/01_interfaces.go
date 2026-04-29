package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Runner interface {
	run()
}

type Jumper interface {
	jump()
}

type Sporter interface { // 接口名后er
	run() // 名称（run）一致，入参（()内）一致，出参（均没有）一致。只有声明，没有实现，具体实现要看具体类型
	// jump() // 加了新的就不算
	Runner // 此时相当于接口嵌套，Runner里面有run()方法
	Jumper // 接口嵌套：此时a既是Person的实例，又是Sporter接口的，又是Runner类型的，又是Jump类型的
}

func (Person) run() { // *Person实现了Sporter接口（*Person实现了Sporter的方法）
	fmt.Println("Run ~~~")
}

func (*Person) jump() {
	fmt.Println("Jump ~~~")
}

// 匿名定义接口（全局变量），不建议
var x interface {
	run()
}

func main() {
	a := new(Person) // 定义Person实例，拿到指针。接口悬空时调用操作非常危险，正常先要判断是否为空
	a.run()          // 调用
	a.jump()

	var s Sporter
	fmt.Printf("%v\n", s)

	s = a // 虽然类型不一致但是相等，实例和指针都是可以访问方法的。可以这样理解，你实现了接口，即符合规范，就相当于又是这个类型，也是另一个类型
	// s2 = Sporter // 不能这样写，接口是规范，它的变量应该符合该接口的实例或实例的指针
	// a 是 Person类型的，此例中确切是*Person
	// 由于*Person实现了Sporter接口的所有方法，*Person实现了Sporter接口
	// a是Person的实例的指针，不准确地说，a即使Person类型的，也是Sporter类型的

	s.run()
	s.jump() // 现在可以调用，因为jump()已经在接口规范里

	var r Runner
	// r = &a
	r = a
	r.run()
	// r.jump // 做不到，不在作用域内

	// 局部变量空接口类型，相当于任何类型（any）
	var y interface{} // 当前悬空nil
	y = 123           // int类型相当于实现了该接口，所以能够修改
	fmt.Printf("%T %[1]v\n", y)
	y = "abc"
	fmt.Printf("%T %[1]v\n", y)

	// x.run()	// 空接口没有方法，.方法直接报错

	// 关系
	b := 123
	y = b
	c := "abc"
	y = c
	// c = b  // 不能得到该结论

	// 切片
	// []any []interface{}	// 表示可以接收任意类型的实例，用的较少，因为Go语言习惯类型是明确的
	z := []interface{}{ // 前是类型，后花括号是字面量需要的
		1, 1.2, "abcd", []int{1, 2}, [2]int{1, 2}, map[int]string{}, // 可以任意类型
	}
	fmt.Printf("%T %[1]v\n", z)
}
