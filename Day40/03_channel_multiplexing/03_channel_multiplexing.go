package main

import (
	"fmt"
	"time"
)

func main() {
	count := make(chan int, 4)
	fin := make(chan struct{})

	// 定时控制
	//t1 := time.NewTicker(2 * time.Second)
	//t2 := time.NewTicker(time.Second)

	go func() {
		defer func() {
			fin <- struct{}{}
		}()
		for i := 0; i < 10; i++ {
			count <- i + 100
			time.Sleep(time.Second)
		}
	}()
	//LOOP:
	for {
		select { // 该关键字用于监听channel，解决Go的多路复用
		// 监听多路
		case <-fin:
			fmt.Println("结束了")
			//break LOOP // break在select里要使用标签语句打破for循环（因为默认break select），这里如果不打破for循环就一直在监听，也会产生死锁
			//goto END   // 或使用goto
			return // 推荐使用return
		case n := <-count:
			fmt.Println("count=", n)
			//default: // 如果以上通道都不满足，卡住了，default执行
			//	fmt.Println("!!!!!!!!!!!!!!!!!!!")
			//	time.Sleep(time.Second)
		}
	}
	//END:
	//	fmt.Println("main结束")
}
