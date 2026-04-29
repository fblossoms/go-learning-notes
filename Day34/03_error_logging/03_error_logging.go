package main

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = "2006/01/02 15:04:05 -0700"
}

func add(x, y int) int {
	return x + y
}

func div(x, y int) int {
	return x / y
}

func calc(x, y int, fn func(x, y int) int) int {
	return fn(x, y)
}

func main() {
	defer func() {
		err := recover()
		fmt.Println(err)
		switch v := err.(type) {
		case nil:
			log.Debug().Send() // 如果是nil建议就直接放行
		case runtime.Error:
			fmt.Println(v)
			fmt.Println(string(debug.Stack()))
			log.Error().Err(v).Str("stack", string(debug.Stack())).Caller().Send() // 长度过长
		default:
			fmt.Printf("其他错误，类型%T，%[1]v\n", v)
		}
	}()
	fmt.Println(calc(4, 5, add))
	fmt.Println(calc(4, 5, div))
	fmt.Println("=========================================")
	fmt.Println(calc(5, 0, div))
}
