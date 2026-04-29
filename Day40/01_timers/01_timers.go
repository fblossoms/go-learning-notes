package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.NewTimer(2 * time.Second) // 一次性的
	// C指的是Channel相当于make(chan Time, 1)，只能放一个
	fmt.Println(<-t1.C)
	fmt.Println(<-t1.C) // 第二次直接报死锁

	t2 := time.NewTicker(5 * time.Second) // 心跳间隔时间 interval
	for {                                 // 间隔造一个就拿一个
		fmt.Println(<-t2.C)
	}
}
