/*
Go语言注释：
	- 单行：// 内容
	- 区间：/*
			内容
           /*
	- 标记：// TODO 某功能未完成
- 建议使用单行和多行都使用 //
- 每行代码后不需要写分号，写完回车即可
*/

package main

import "fmt"

func main() {
	// 注释
	fmt.Println(min(1, 2)) // 后面写的注释说明改行的作用，结果打印到控制台
	var a int = 100        // Go语言中，变量未使用也会报错，即定义的变量必须使用
	fmt.Println(a)
	a = 300
	/*fmt.Println(a)*/
}
