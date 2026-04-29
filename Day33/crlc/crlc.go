// 子包示例
package crlc

import "fmt"

func Add(x, y int) int {
	fmt.Printf("calc/calc.go Add: %d, %d\n", x, y)
	return x + y
}
