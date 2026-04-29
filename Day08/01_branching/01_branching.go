package main

import "fmt"

func main() {
	var a int = -100
	if a > 0 {
		fmt.Println("a为正数") // 这个满足条件，下面的都不执行
	} else if a == 0 {
		fmt.Println("a为0")
	} else {
		fmt.Println("a为负数")
	}

	var b = 200
	if b == 0 {
		fmt.Println("b为0")
	} else {
		if b > 0 {
			fmt.Println("b为正数")
		} else {
			fmt.Println("b为负数")
		}
	}

	e := 399
	switch /*有个看不见的布尔值，默认为true*/ {
	case e == 0:
		fmt.Println("e等于0")
	case e > 0:
		fmt.Println("e为正数")
	case e < 0:
		fmt.Println("e为负数")
	default:
		fmt.Println("e为虚数")
	}

	const f int = 200
	switch f {
	case 200:
		fmt.Println("f等于200")
	case 4, 8, 16, 24:
		fmt.Println("f等于4")
	}

	switch g := 199; {
	case g == 0:
		fmt.Println("g为0")
	case g > 0:
		fmt.Println("g大于0")
	}

	var score, line int
	if score, line := 92, 96; score > line { // 为当前if重新定义当前if-else自己使用的score和line
		fmt.Println("恭喜你，已过分数线")
	} else {
		fmt.Println("很遗憾，你未过分数线")
	}
	fmt.Printf("score=%v，line=%v\n", score, line)

	var score2, line2 int
	if score2, line2 = 92, 96; score2 > line2 { // 没有为当前if重新定义当前if-else自己使用的score和line，而是赋值
		fmt.Println("恭喜你，已过分数线")
	} else {
		fmt.Println("很遗憾，你未过分数线")
	}
	fmt.Printf("score=%v，line=%v\n", score2, line2)
}
