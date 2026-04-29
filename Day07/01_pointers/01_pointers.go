package main

import "fmt"

func main() {
	var a = 100
	b := &a
	fmt.Println(&a, b, a, *b)
	fmt.Println(&a == b, a == *b) // 注意：Go中不能使用b++、b--这种写法，因为类型为*int（指针类型），防止内存越界危险

	d := a
	fmt.Println(&d, &a) // 存的位置不一样，因为用了其他的内存地址存a的地址
	fmt.Println(d == a, &d == &a)

	var c *int // 初始化，则给零值，指针类型对应的零值为nil
	fmt.Println(c)
	
}
