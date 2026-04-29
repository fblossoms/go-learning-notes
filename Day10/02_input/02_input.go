package main

import "fmt"

func main() {
	var n int
	var err error             // error类型，内建的接口类型，零值为nil
	var a, b string           // 零值，输入类型与定义类型一致，因为scan行做了力所能及的类型转换
	n, err = fmt.Scan(&a, &b) // 类型为指针，取地址，通过地址操作a和b，可以提升性能。
	// 多个值利用空白分割，换行符也是空白，即通过回车输入下一值。输入数大于变量数，只取变量数的值
	// n表达接收到控制台传进来的几个值，err表示是否有报错，输入类型与定义类型不一致导致无法进行类型转换就会报错
	fmt.Println(n, err, a, b)

	var name string
	var age int
	fmt.Println("请按顺序输入你的姓名和年龄")
	n, err = fmt.Scanf("%s %d", &name, &age) // \n换行符不再作为空白
	// n, err = fmt.Scanf("%s,%d", &name, &age) // 格式化输入会把逗号（,）当作字符串和后面的age一起赋给%s。该情况易出bug
	// n, err = fmt.Scanf("%d，%s", &age, &name) // 可调整顺序
	fmt.Println(err, name, age)

	// 生产级写法：建议一般使用Scan和Scanln即可，Scanf易产生bug
	var input string
	fmt.Scan(&input) // 直接用变量接收整个字符串，因为永远不要相信客户端表单提交的数据，后续自行通过限定进行分割类型
	// 或者分行写，一行传入一个值
}
