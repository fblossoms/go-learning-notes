package main

import (
	"fmt"
	"math" //本地源代码：GOROOT下src目录就是go环境安装时候
)

func main() {
	// int
	var a, b = 1.44, -2.66
	fmt.Println(int(a), int(b)) // 正数时向下取整，负数时向上取整

	// 除法
	fmt.Println(1/2, 12/7, -1/2, -12/7) // 正数时向下取整，负数时向上取整
	fmt.Println(1.0 / 2)                // 除法取整仅限整数相除使用

	// math.Ceil
	fmt.Println(math.Ceil(2.32), math.Ceil(-2.32)) // 向上取整

	// math.Floor
	fmt.Println(math.Floor(2.32), math.Floor(-2.32)) // 向下取整

	// math.Round
	fmt.Println(math.Round(2.32), math.Round(-2.32))                               // 四舍五入
	fmt.Println(math.RoundToEven(0.5), math.RoundToEven(1.5), math.RoundToEven(3)) // 取整到最近偶数，碰到奇数则不取整

	fmt.Println(math.Abs(-2.1)) // 绝对值
	fmt.Println(math.Pi)        // 圆周率
	fmt.Println(math.E)         // 自然常数

	fmt.Println(math.Max(2.45, 6.65)) // 取最大值
	fmt.Println(math.Min(2.45, 6.65)) // 取最小值

	fmt.Println(math.NaN()) // not a number（字面意思：不是一个数字）类型为float64，其实也算是数字，无法与其他数字进行比较

	fmt.Println(math.Pow(2, 3))                 // 以x为底的y次幂
	fmt.Println(math.Sqrt(2), math.Pow(2, 0.5)) // x开方

	fmt.Println(5%2, math.Mod(5, 2)) // 取模（取余）

	// 1.21版本+ 内建函数（无需调包）
	fmt.Println(max(1, 2, 3, 4.4, 9.1))                // 同类型比较大小（整型和浮点型可通过隐式类型转换使用该比较语法）
	fmt.Printf("%v\n", min("abc", "ABC", "qmn", "xyz", "\n")) // 字符串也可以用来比较。底层为排序算法
}
