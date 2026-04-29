package main

import "fmt"

func main() {
	var a rune = 'a'                                  // 字符类型为rune，本质为int32
	fmt.Printf("%T, %[1]d, %[1]c, %[1]x, %[1]q\n", a) // 可以用%d打印，因为是int32

	a = '对' // 正确，因为int32四字节，能够放下中文
	fmt.Printf("%T, %[1]d, %[1]c, %[1]x, %[1]q\n", a)

	var e byte = 0x61
	fmt.Printf("%T, %[1]d, %[1]c, %[1]x, %[1]q\n", e)
	e = '\x61' // 能运行，涉及隐式类型转换。有风险，因为类型不匹配（byte和rune）
	fmt.Printf("%T, %[1]d, %[1]c, %[1]x, %[1]q\n", e)
	e = 'a' // 内存中的值相同，都是整数
	fmt.Printf("%T, %[1]d, %[1]c, %[1]x, %[1]q\n", e)
	// e = '对' // 中文字符超界（超过1字节）

	var f = 27979 // 默认为int，但会找到对应字符。int（64）比int32更占空间
	fmt.Printf("%T, %[1]d, %[1]c, %[1]x, %[1]q\n", f)
}
