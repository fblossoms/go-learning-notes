package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 引入 等待组
	var wg sync.WaitGroup // 结构体，倒计数器，计数为0则不用再wait
	count := 5
	wg.Add(count) // 添加需要等待的协程数量，当前为 1 组

	fmt.Println("协程数：", runtime.NumGoroutine()) // 查看当前协程数量
	fmt.Println("4 main start")

	go func() {
		defer wg.Done() // 儿子协程也要减
		fmt.Println("我是被主协程创建出来的协程，我准备创建其他子协程")
		for i := 0; i < count-1; i++ {
			go func(i int) {
				defer wg.Done()
				defer fmt.Printf("我是被创建的孙子协程 %d, byebye\n", i)
				fmt.Printf("我是被创建的孙子协程 %d\n", i)
				time.Sleep(5 * time.Second)
				// fmt.Printf("我是被创建的孙子协程 %d, byebye\n", i)
			}(i)
		}
	}()

	fmt.Println("协程数：", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("5 main end")
	fmt.Println("协程数：", runtime.NumGoroutine())
}
