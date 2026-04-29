package main

import (
	"fmt"
	"time"
)

type Person struct{} // 无数据成员的结构体，实例化不占内存，数据占用空间0个字节

func main() {
	//m := Person{} // 等价于struct{}{}

	c := make(chan struct{}) // 解除阻塞（0byte）

	d := make(chan bool) // 可以传bool（1byte）
	e := make(chan int)  // 可以传int（8byte）
	fmt.Printf("%T len=%d, cap=%d\n", c, len(c), cap(c))
	fmt.Println()

	go func() {
		time.Sleep(5 * time.Second)
		c <- struct{}{} // 业务上没有意义，但是打入空结构体实际上是为了解除阻塞
		close(c)        // 接触阻塞，但是默认零值
		d <- true
		e <- 100
	}()
	//x, ok := <-c
	//fmt.Println("我终于等到你了", x, ok)
	fmt.Println("我终于等到你了", <-c)
	fmt.Println("我终于等到你了", <-d)
	fmt.Println("我终于等到你了", <-e)
}
