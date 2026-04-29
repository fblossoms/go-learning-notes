package main

import (
	"errors"
	"fmt"
	"runtime"
)

// 定义错误类型，首字母大写为包外可用
var ErrDivisionByZero = errors.New("除零异常") // 既是*errors.errorString类型，又是error接口类型

// 模拟错误：除数为0
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero // errors包New构建error类型的值
	}
	return a / b, nil // 没错自然返回零值
}

// 解释方法拓展（既是又是）
type MyInt2 int // 此时无法实现error功能

func (*MyInt2) Error() string {
	return "我乐意返回这个字符串"
}

func foo() (err error) {
	var x MyInt2 = 200
	err = &x // 相当于 var err error = &x
	return err
}

// panic
func div2(a, b int) (int, error) {
	defer fmt.Println("1 进入div")
	defer fmt.Println(2, a, b)
	r := a / b
	fmt.Println(3, r)
	return r, nil
}

// recover
func div3(a, b int) (int, error) {
	defer func() {
		// recover() // 处理报错，恢复正常
		err := recover()
		fmt.Printf("%+v, %[1]T\n", err) // 接收报错原因
		if v, ok := err.(error); ok {   // 通过断言拓展类型
			fmt.Println(v.Error())
		} else {
			fmt.Println("管不了")
		}
	}()
	defer fmt.Println("1 进入div")
	defer fmt.Println(2, a, b)
	r := a / b
	fmt.Println(3, r)
	return r, nil
}

// recover的switch写法
func div5(a, b int) (int, error) {
	defer func() {
		err := recover()
		fmt.Printf("%T %[1]v\n", err)
		switch v := err.(type) {
		case nil:
			fmt.Println("没有错误", nil)
		case runtime.Error: // 小类型
			fmt.Println("运行时错误", v)
		case error: // 大类型，优先走，若写在runtime.Error前，则 runtime.Error永远不可到达
			fmt.Println("我是error接口类型的", v.Error())
		default:
			fmt.Println("其他错误类型，我管不了")
		}
	}()
	r := a / b
	return r, nil
}

// recover2
func div4(a, b int) (int, error) {
	defer fmt.Println("1 进入div")
	defer fmt.Println(2, a, b)
	r := a / b
	fmt.Println(3, r)
	return r, nil
}

func main() {
	if r, err := div(5, 0); err != nil { // 根据位置得到函数的返回值
		fmt.Println("错误：", err.Error()) // 默认调用 err类型的error接口类型的Error()方法
		fmt.Println("错误：", err)
	} else {
		fmt.Println("没有错误", r)
	}
	fmt.Println("=================================================================")

	// panic
	if r, err := div2(5, 2); err != nil {
		fmt.Println("4 错误：", err) // 拿不到返回值，则无法进入语句，从而报错，进而导致main函数崩
	} else {
		fmt.Println("没有错误", r)
	}
	fmt.Println("=================================================================")

	// recover
	if r, err := div3(5, 0); err != nil {
		fmt.Println("4 错误：", err) // 错误已经在func中被recover处理了，r和err均为零值，不会走该分支
	} else {
		fmt.Println("没有错误", r)
	}
	fmt.Println("进程正常结束") // 若div3报错，导致main函数无法解决而崩掉，不会看到该打印
	fmt.Println("=================================================================")

	// recover2：recover在main中
	defer func() {
		err := recover() // 对函数爆出来的panic进行处理，但是没有div4的返回值，if后续将不执行
		// 如果没有错也会执行，此时err为零值nil
		fmt.Printf("6 在main中管panic, %v\n", err)
	}()
	if r, err := div4(5, 0); err != nil {
		fmt.Println("4 错误：", err)
	} else {
		fmt.Println("没有错误", r)
	}
	fmt.Println("7 进程正常结束")
	fmt.Println("=================================================================")

	// recover的switch写法
	if r, err := div5(25, 4); err != nil {
		fmt.Println("错误", err)
	} else {
		fmt.Println("没有错误", r)
	}
}
