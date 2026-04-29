package main // 同一个目录里，包名必须一致

import (
	"fmt"
	"os"
)

const a = 100 // 包内全局使用。全局常量不可以被取地址（全局变量可用被取地址）

var (
	b = 222
	c = 333
)

func Test1() { // 函数名大写后，包外可用。相当于导出。使用时按照 包名.函数的格式引用，例如当前的main.Test
	// Test()
}

func showBValue() int {
	fmt.Printf("%v %p\n", b, &b)
	return b
}

func main() {
	// 1、语句块作用域
	s := []int{1, 3, 5}
	for i, v := range s {
		fmt.Println(i, v) // i和v在for块中可见
	}
	// fmt.Println(i, v) // 错误，在for外不可见

	if f, err := os.Open("o:/t.txt"); err != nil {
		fmt.Println(f, err) // 可见
	}
	// fmt.Println(f, err)	// 错误，不可见

	// 2、显式的块作用域
	{
		// 块作用域
		const a = 100
		var b = 200
		c := 300
		fmt.Println(a, b, c) // 可见
	}
	// fmt.Println(a, b, c) // 错误，不可见

	// 3、universe
	const c = iota // 内建包，将被加载到全局空间（当前程序的所有地方都可以直接使用并且不需要写包名）

	// 4、函数（当前为main）内可用
	// const a = 100

	// Test1() // 运行时需要编译所有相关文件

	a := 1 // 优先使用局部变量的值，并且局部变量无法影响当前作用域之外的值
	fmt.Println(a)

	// 观察b
	fmt.Println("第1次", b, &b)        // 由于当前函数中没有定义b，所以使用全局变量b
	b = 200                          // 修改全局变量的值
	fmt.Println("第2次", b, &b)        // 由于全局变量b的值被修改，所以值为修改后的值，地址为全局变量b的地址
	b := 300                         // 定义了当前函数内自己的b
	fmt.Println("第3次", b, &b)        // 由于当前函数定义了b，所以使用局部变量b
	fmt.Println("第4次", showBValue()) // 由于showBValue()这个函数里访问的b，只可能是包级（全局）b，它看不到 main()里的局部b

}
