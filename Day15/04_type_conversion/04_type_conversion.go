package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Weekday int

const (
	Sun Weekday = iota // 上面定义了一个新类型Weekday，用无类型字面常量iota可以转换成Weekday
	Mon
	Tue
)

func testWeekday(w Weekday) {
	fmt.Println("星期几 =", w)

}

func main() {
	// 类型转换
	var a int8 = -1
	var b = uint8(a) // 需要手动进行类型转换
	fmt.Println(a, b)

	var c = 3.14 // 变量才可以进行手动类型转换，字面常量不能手动类型转换
	fmt.Println(int(c))

	m := 'a'
	fmt.Printf("%c, %[1]T\n", m+1) // int32 + int -> int32，物业类型字面常量做了隐式类型转换（int32(1)）
	n := 1
	fmt.Printf("%c, %[1]T\n", m+int32(n)) // 变量需要手动做类型转换

	// 类型别名
	var a2 byte = 'a' // byte就是uint8，byte是uint8的别名
	var b2 uint8 = 98
	fmt.Printf("%T, %[1]c, %[1]d", a2+b2)

	type mybyte uint8 // 没有等号，相当于前者基于后者的新类型，本质上是后者的类型。但是Go中任务不是同一种类型，想要计算需要做手动类型转换
	// 虽然只是基于，但是可以在后续定义新方法，更加灵活方便
	type Mybyte = uint8 // 有等号就是等价，才可以做相同类型计算
	var c2 mybyte = 99
	var d2 Mybyte = 99
	fmt.Printf("%T, %[1]c, %[1]d\n", a2+uint8(c2))
	fmt.Printf("%T, %[1]c, %[1]d\n", b2+d2)

	// 转换字符串，无类型字面常量也能转，解释为代码点
	fmt.Println(string(rune(97)), string(rune(27979)))
	// 转换成数字字符
	fmt.Printf("%T, %[1]v\n", strconv.Itoa(97))

	// strconv库
	fmt.Println(strconv.ParseBool("1")) // Go中不会把1当做bool，但是这个库可以
	fmt.Println(strconv.ParseBool(strings.ToLower("TruE")))
	fmt.Println(strconv.ParseFloat("12.62", 64)) // bitSize为预期转换类型

	// 常用操作
	// string -> []byte、[]rune
	s1 := "abcd"
	s2 := "测试"
	fmt.Println(len(s1), len(s2))
	fmt.Println([]byte(s1), []byte(s2)) // 按字节分
	fmt.Println([]rune(s1), []rune(s2)) // 按字符分，1个字符4个字节

	// []byte、[]rune -> string
	fmt.Println(string([]byte{97, 0x62}))                         // 不能加超界（255）的字符
	fmt.Println(string([]rune{97, 27979}))                        // 可以中文
	fmt.Println(string([]byte{97, 0x41, 0x30, 0xe6, 0xb5, 0x8b})) // 可以识别出中文

	// 枚举：让你写的东西，在枚举定义的常量范围之内，如果不在，就给你报错
	testWeekday(0)
	var l = 1
	testWeekday(Weekday(l)) // 使用变量可以让编译帮做类型检查，类型不一致，需要用手动类型转换，出无类型字面常量外
	testWeekday(Sun)
}
