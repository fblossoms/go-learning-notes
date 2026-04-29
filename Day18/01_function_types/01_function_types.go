package main

import (
	"fmt"
	"os"
)

func fn1() { // 函数名首字母小写为包内函数，仅可在包内使用，包外看不见
	return // 只写return表示仅仅返回，意味着函数执行完成了，结束了，栈帧消亡。不写也默认会隐含return（如果没值）
	// return后不可使用a++语法，因为a++为先把a使用，再+1，不符合函数要求。但可以使用 a + 1扩展，其他语言可以使用return ++a语法，先 +1再使用
}
func fn2(i int) int { // 这里 )后的 int没有给名字，称为定义匿名返回值。定义时要么全匿名要么全具名，不建议混合
	a := 1000      // 函数内定义的变量称为局部变量，作用只在该函数体代码块生效
	return 100 - a // 定义返回值列表时，return的必须写值，且类型要与定义的返回值类型一致
}
func fn3(j int) (r int) { // 这里的(r int)为定义具名返回值。除没有返回值和单个返回值类型之外，其他多返回值或具名返回值必须使用括号
	// return // 定义具名返回值函数可以写return相当于reruen r，返回为该对应类型零值
	// return r + 100 // 此时相当于返回0 + 100即100
	r = 200 // 可以赋值，但注意不要重新定义，因为定义具名返回值是已经定义
	return r + 100
}

// 多返回值
func fn4() (int, int) { // 定义返回值列表时使用()，多个相同类型必须重复定义，不能缺省
	// return 1, -1 // 必须一一对应返回值列表的个数和类型写出返回值
	a, b := 1, -1
	a, b = b, a+b // 交换变在量
	return a, b
}

func fn5() (x int, y bool) {
	return // 具名函数可以写return，因为定义时已经初始化。匿名函数多个返回值不能直接写return
}

// 接口类型若没使用具名函数必须写值
func fn6() error {
	return nil // 内建类型，接口类型，错误接口类型，接口类型零值都是nil
}

func fn7() (err error) {
	// return nil
	t, err := os.Open("o:/test1.txt")
	fmt.Println(t)
	fmt.Println(err, "####", err == nil) // err.String() 描述，有描述，有错误，错误就不是nil
	return
}

func fn8() (err error) {
	if t, err := os.Open("o:/t.txt"); err != nil { // 这里的err属于重新定义
		fmt.Println(err, "%%%%")
		fmt.Println(t)
		return err // 此时这里的返回值不能省，因为这里的err在时if定义的局部变量，与fn8函数无关。只有具名返回值才能省略return返回值
	}
	// else
	return
}
func main() { // 不允许在函数内使用func定义新函数，想定义只能定义到外部
	fmt.Printf("%T\n", fn1)
	fmt.Printf("%T\n", fn2)
	fmt.Printf("%T, %d\n", fn3, fn3(1))
	// 打印结果
	// func()
	// func(int) int	// 入参类型为int，出参类型也为int
	// func(int) int	// fn1和fn2类型一致。同类型函数即入参和出参类型一致、个数一致

	fn1() // 函数名后面加括号，表示调用这个函数
	// t := fn1() // 没有返回值的函数不能赋给其他变量，同理，不能使用该值

	fn2(1)              // 该函数有返回值，但是没使用（没有赋予给其他变量），函数依然执行，但是没被使用
	fmt.Println(fn2(1)) // 此时返回值赋给了fmt包的Println函数进行使用

	a, b := fn4() // 多返回值必须使用对应个数的标识符来接收
	fmt.Println(a, b)

	fmt.Println(fn8())
}
