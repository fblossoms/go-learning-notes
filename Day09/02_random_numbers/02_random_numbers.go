package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 兼容1.8前
	//os.Setenv("GODEBUG", "randautoseed=0") // 设置关闭随机种子，此时起始值、公式一样，五位随机数依次恒为17791

	// 兼容1.6前
	//var seed int64 = 1
	//src := rand.NewSource(seed)
	//r := rand.New(src) // r为随机数生成器

	// 工程写法：用处例如负载均衡
	//var seed2 int64 = time.Now().UnixNano() // 当前时间的纳秒时间戳，返回为整数int64
	//src2 := rand.NewSource(seed2)
	//r2 := rand.New(src2)
	r2 := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 10; i++ {
		fmt.Println(i, r2.Intn(10), rand.Intn(10)) // [0, 9]，[0, 10)前包后不包
		fmt.Println(i, rand.Intn(100)+100)         // 加偏移量[100, 200)
	}

}
