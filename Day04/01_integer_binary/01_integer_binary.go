package main

import "fmt"

func main() {
	var a = 1
	var b = 2.1
	fmt.Println(a, b)            // Println为打印后换行
	fmt.Printf("%T, %T\n", a, b) // Printf能够按照格式输出，%T为取类型，依次取类型

	//fmt.Println(a + b) // 类型不同的变量标识符无法相加（变量无法进行隐式类型转换）
	fmt.Printf("%T\n", 1+2.1) // 无类型字面常量可以使用隐式类型转换，转换方向朝精度高的转int => float64
	//fmt.Println(a + 2.1)	// 无法运算，需要自行截断float64类型的2.1为int与a的类型相同才可运算
	//fmt.Println(a + int(2.1))	// 字符常量无法使用int()进行强制类型转换，需要使用变量
	fmt.Println(a + int(b)) // 变量可以进行强制类型转换

	fmt.Printf("%T %[1]v\n", 1+2.1) // 等价于%[1]T %[2]v，由于2没有值对应所以丢失，可以手动改索引。%v为取值
	fmt.Printf("%T %[1]v", 1+2.2)

	var c int64 = 100 // 100为无字符字面常量，默认为int，隐式转换成了int64
	t := 100
	//var d int64 = t // t类型推断为int，而t变量，无法隐式转换
	var d int64 = int64(t) // 变量可以强制类型转换
	fmt.Printf("\n%T, %[1]v", c)
	fmt.Printf("\n%T, %[1]v", d)

	var e int64 = 0x32 // 十六进制，打印十进制对应为50
	fmt.Println("\n", e)

}
