package main

import "fmt"

func main() {
	// iota 只能用于常量初始化
	const a1 = iota
	// var b = iota
	fmt.Println(a1)

	// iota 可以连续定义
	const a2 = iota
	const b2 = iota
	const c2 = iota
	fmt.Println(a2, b2, c2)

	// iota 可以批量定义，但会递增
	const (
		a3 = iota
		b3 = iota
		c3 = iota
	)
	// 批量定义时右边计算公式一致，可以只写一遍（在第一个）
	const (
		a4 = iota
		b4
		c4
	)

	const (
		a = iota
		b
		c
		_ // _会废掉当前值3
		d
		e = 100 // 规则改变
		f       // 按照最新公式规则为准
	)
	const (
		_ = 1
		m
		n
		a0 = iota // 只要const里出现了iota，此时第一个iota的值相当于从行索引值开始
		b0
		c0
		_
		d0
		e0 = 100 + iota // iota可以正常参与计算
		f0
	)
}
