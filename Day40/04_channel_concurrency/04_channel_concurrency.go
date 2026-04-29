// TODO 需求：有一个全局数count，初始为0。编写一个函数inc，能够对count增加10万次。执行5次inc函数，请问最终的count值是多少

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var count int64 = 0
var mx sync.Mutex // 排他（互斥锁），第二值为倒计时器，其他协程碰到会阻塞（谁碰谁阻塞）。使用完后按照队列顺序依次进入使用
var ch = make(chan int64, 1)

func inc() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		// atomic 原子性，原子不可分割
		// count++ // 非原子操作，每次操作完可能会被打断，打断后需要重新操作相同数据（多个协程对数据同时操作，覆盖等），会出现安全问题，所以需要进行隔离（排他，使用原子）
		//count = count + 1

		//atomic.AddInt64(&count, 1)

		//mx.Lock()   // 拿不到锁，协程阻塞，如果不阻塞，说明你是唯一拿到锁的，用有绝对的使用权，使用完后注意要解锁
		//count++     // 计算密集型，不可并行，疯狂使用CPU资源
		//mx.Unlock() // 解锁

		// 通道并发，通道内置有锁
		t := <-ch
		t++
		ch <- t
	}
}

func main() {
	start := time.Now()

	// 串行 大概500微秒
	//for i := 0; i < 5; i++ {
	//	inc()
	//}

	// 并行 大概4000微秒，并发，不安全，生产上禁止使用
	// 并行atomic 大概6000微秒，增加了时间
	// 并行atomic排他（互斥）锁	大概8000微秒，这里相当于五车道公路汇聚到一车道的收费站，排队（谁在排队谁阻塞）依次进入收费站，过完收费站后又是五车道。不适合并发，但安全
	// 并行channel 大概100000微秒，排他锁越多越影响速度，没把channel的用处发挥出来
	wg.Add(1)
	for i := 0; i < 1; i++ {
		go inc()
	}
	ch <- 0 // 使用并发通道前需要初始化

	fmt.Printf("协程数 %d\n", runtime.NumGoroutine())
	wg.Wait()

	fmt.Printf("协程数 %d\n", runtime.NumGoroutine())
	fmt.Printf("执行时长为 %d微秒\n", time.Since(start).Microseconds())
	//fmt.Printf("count = %d\n", count)
	fmt.Printf("<-ch = %d\n", <-ch)
}
