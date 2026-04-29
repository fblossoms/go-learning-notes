package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 单向通道：防止传参误操作
func producer(ch chan<- int) {
	for {
		ch <- rand.Intn(20)
		time.Sleep(time.Second)
	}
}
func consumer(ch <-chan int) {
	for {
		t := <-ch
		fmt.Printf("从队列中消费一个数据 %v\n", t)
	}
}
func main() {
	// var c1 chan int // 不能用，零值不可用，nil通道不可用，底层根本就没有容器
	// fmt.Printf("%v, len=%d, cap=%d\n", c1, len(c1), cap(c1))

	c2 := make(chan int, 8) // 缓冲通道，容量8
	c3 := make(chan int, 0) // 一个元素都放不下，底层数据结构，unbuffered 非缓冲通道（同步通道）。0可以省略不写
	fmt.Printf("%v, len=%d, cap=%d\n", c2, len(c2), cap(c2))
	fmt.Printf("%v, len=%d, cap=%d\n", c3, len(c3), cap(c3))

	// c3缓冲通道举例
	go func() {
		time.Sleep(3 * time.Second)
		c3 <- 1000
	}()
	// send
	x := <-c3
	fmt.Println(x)

	// c2缓冲通道举例
	c2 <- 111
	c2 <- 222
	c2 <- 333
	y := <-c2
	fmt.Println(y)

	// c4单向通道，没有使用的意义
	c4 := make(chan<- int, 1) // 只进不出
	fmt.Printf("%v, len=%d, cap=%d\n", c4, len(c4), cap(c4))
	c4 <- 444

	// 单向通道：调用
	c5 := make(chan int, 1)
	go consumer(c5)
	go producer(c5)
}
