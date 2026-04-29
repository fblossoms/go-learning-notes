package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func add(x, y int, w *sync.WaitGroup) int {
	var c int
	defer w.Done() // WaitGroup计数器减1
	defer fmt.Printf("1 return c = %d\n", c)
	defer func() { fmt.Printf("2 return c = %d\n", c) }()
	fmt.Printf("3 add called, x = %d, y = %d\n", x, y)
	c = x + y
	time.Sleep(3 * time.Second)
	return c
}

func main() {
	// 引入 等待组
	var wg sync.WaitGroup // 结构体，倒计数器，计数为0则不用再wait
	wg.Add(1)             // 添加需要等待的协程数量，当前为 1 组

	fmt.Println("4 main start")
	fmt.Println("协程数：", runtime.NumGoroutine()) // 查看当前协程数量
	// add(4, 5) // 普通函数调用，返回才能走到下一行
	// time.Sleep(time.Second) // 手动等待

	go add(4, 5, &wg) // 创建协程
	fmt.Println("协程数：", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("5 main end")
	fmt.Println("协程数：", runtime.NumGoroutine())
}
