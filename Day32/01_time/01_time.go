package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now() // 返回当前时间
	fmt.Printf("%T %+[1]v\n", t1)
	// +0800 CST 为用户机器的时间，当前是东八区时间（北京时间），东区为+号，西区为-号
	fmt.Printf("%T %#[1]v\n", t1)

	t2 := t1.UTC() // 世界协调时
	fmt.Printf("%T %+[1]v\n", t2)

	//time.NewTicker() // 时间间隔器
	//time.NewTimer()  // 计时器

	// 时间格式化，Go不支持C风格格式符
	fmt.Println(t1.Format("01 02 15 04 05 2006 -0700"))
	// 06/01/02 03:04:05
	// 06代表年（前面加20为公元纪年制），01代表月，02代表日
	// 03代表12小时制小时（15为24小时制），04代表分钟，05代表秒，-0700代表西七区
	// 忘记格式就看time文档

	// 通常写法，输出时建议带上时区
	fmt.Println(t1.Format("2006/01/02 15:04:05.000000 -0700"))
	// .0000000超过7位的0为无效数字，精度达不到，就按照位数补齐

	// string解析为time.Time时间对象
	s1 := "2008/10/07 13:06:56.1234567 +0800"
	t3, err := time.Parse("2006/01/02 15:04:05.0000000 -0700", s1) // 解析时小数部分的位数要对应
	// 不带时区就默认零时区，所以一定要加上，否则会影响后续操作，或者使用PaserInLocation
	if err != nil {
		panic(err)
	}
	fmt.Println(t3)

	// 提取时间成分
	fmt.Println(t3.Year(), t3.Month(), t3.Day(), int(t3.Month())) // 月份用到了枚举
	fmt.Println(t3.Date())                                        // Date()一次取得年月日
	fmt.Println(t3.Hour(), t3.Minute(), t3.Second())
	fmt.Println(t1.YearDay())
	fmt.Println(t1.Weekday(), int(t1.Weekday()))

	// 时间戳
	fmt.Println(t1.Unix(), t1.UnixMilli(), t1.UnixMicro(), t1.UnixNano())
	// 秒、毫秒、微秒、纳秒

	// 时间戳 -> time.Time对象
	t5 := time.Unix(1773754356, 0)
	fmt.Println(t5)
	fmt.Println(t5.Local(), t5.UTC())

	// 构造当前城市所在时区城市的时间
	tz, _ := time.LoadLocation("Asia/Shanghai")
	t4 := time.Date(2008, 7, 12, 10, 5, 6, 0, tz)
	fmt.Println(t4)

	// 时间运算
	delta := t1.Sub(t3)
	fmt.Printf("%T, %[1]v\n", delta)
	fmt.Printf("%T, %[1]v\n", delta.Hours()/24/365) // 相差年数
	fmt.Printf("%T, %[1]v\n", delta.Seconds())      // 相差秒数

	t8 := t1.Add(2 * time.Second) // 2是纳秒大小，秒要纳秒乘10的9次方
	delta = t8.Sub(t1)
	fmt.Println(delta.Seconds())
	fmt.Println(t8.After(t1)) // 判断t8是否在t1之后（）

	fmt.Println(time.Since(t8)) // 当前的时间 - t8
}
