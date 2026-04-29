package main // 包，包名自定义，当前为main

import "fmt" //	外包的包名，当前为引用一个叫fmt的包

var b int = 20 // b为小写，为当前包（main）的全局变量
var C int = 30 // C为大写，可以在其他包调用这个全局变量

func main() {
	var a int = 10
	fmt.Println(a, b, C) // 使用外包时，首字母大写，这里P为大写
}
