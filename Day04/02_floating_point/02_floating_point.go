package main

import (
	"fmt"
	"math"
)

func main() {
	t := 2.3
	fmt.Printf("1	%f\n", t)        // 精度默认为6，小数点后6位
	fmt.Printf("2	%.2f\n", t)      // 调整精度：小数点后2位
	fmt.Printf("3	%2f\n", t)       // 2为宽度，由于没有设置精度，精度默认为6，此时6比2大，为了保持消息完整，以消息为主
	fmt.Printf("4	|%10f|\n", t)    // 此时6比10小，能够体现宽度
	fmt.Printf("5	|%-10f|\n", t)   // 加了-号为左对其
	fmt.Printf("6	|%-10.3f|\n", t) // 可以混合使用

	fmt.Println(
		math.MaxFloat64, math.MaxFloat32,
		math.MaxInt64, math.MaxInt32, math.MaxInt,
		math.MaxUint8, math.MaxInt8, // 最后一个的后面也要加逗号,
	)
}
