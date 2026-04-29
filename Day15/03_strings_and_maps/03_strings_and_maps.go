package main

import (
	"fmt"
	"strings"
)

func main() {
	var a = 'a' // 类型为rune（int32）整数 1个字符，4个字节，a字符4个字节（00 00 00 61）
	fmt.Printf("%T, %[1]d, %[1]c, %[1]x, %[1]q\n", a)

	var b = "a" // 类型为string 线性数据结构体	a字符1个字节
	fmt.Printf("%T, %[1]s, %[1]x, %[1]q\n", b)

	// 切片和子串
	var c = "Guangxi广西"
	fmt.Printf("%s, %s, %s\n", string(c[0]), c[0:len(c)-1], c[:])
	// 使用索引时面对只有一个字符的情况是，需要做显示类型转换；或把%s改成%c（取字符）
	fmt.Println(&c) // 不可以取到特定索引的地址

	// 遍历字符串的字符，按字节遍历
	for i := 0; i < len(c); i++ {
		fmt.Printf("%d: %T, %c\n", i, c[i], (c[i]))
	}

	// for range，按字符首字节遍历
	for i, v := range c {
		fmt.Printf("%d: %T, %c, %v, %[4]T, %[4]c\n", i, c[i], (c[i]), v)
	}

	// 字符串拼接，返回的是全新字符串，有由原先字符串拼接而来。因为字符串“只读，内存不可变”
	// 直接俄相加
	var d = "abcd" + "efg"
	fmt.Println(d)
	fmt.Println('广', "西"[0]) // 索引返回的是3个字节的对用的索引的字节的编码表对应值

	// Sprintf：格式化输出，更灵活，推荐
	e := fmt.Sprintf("%s%s---%s", "abcd", "efg", "hijk")
	fmt.Println(e)

	// strings库，参数1为字符串切片，参数2为分隔符
	f := strings.Join([]string{"abcd", "efg", "hijk"}, "!")
	fmt.Println(f)

	// 多次拼接builder，性能高，用于频繁拼接
	//var builder = strings.Builder{}
	var builder strings.Builder             // 初始化未赋值，默认零值。类似的有bytes
	builder.Write([]byte{97, 98, 99, 0x41}) // 参数为字节切片，注意范围只有1字节（0x00~0x80，0~127）
	builder.WriteByte(122)                  // 只追加一个字节
	builder.WriteRune('c')                  // 追加一个字符
	builder.WriteString(c)                  // 追加一个字符串
	fmt.Println(builder.String())           // .String作用是返回操作完的值

	// 从字符串里找字串，返回对应索引，找不到就返回-1。效率都不高，时间复杂度为O(n)，该用则用，慎用
	fmt.Println(strings.Index(c, "广西"))      // 从正方向找，注意循序不能换，且不能分散在两个字串
	fmt.Println(strings.LastIndex(c, "广西"))  // 从反方向找，但返回的依然是正索引
	fmt.Println(strings.IndexAny(c, "广西南宁")) // 参数2为字符集，返回能找到的字符的最小索引
	fmt.Println(strings.IndexByte(c, 97))    // 找字节对应的编码表的字符，返回值为索引
	fmt.Println(strings.IndexRune(c, 'a'))   // 根据字符类型查找，返回值为索引
	fmt.Println(strings.Count(c, "广西"))      // 对应字符串出现过的次数，返回值为整型
	fmt.Println(strings.Contains(c, "广西"))   // 字符串（参数1）内是否包含字串（参数2），返回类型为bool。本质是Index

	// 大小写转换，用于查询数据库等
	fmt.Println(strings.ToUpper(c)) // 全部大写
	fmt.Println(strings.ToLower(c)) // 全部小写

	// 检查是否以特定字符串开头或结尾，返回类型为bool。效率高，时间复杂度O(1)
	fmt.Println(strings.HasPrefix(c, "ab")) // 判断是否在开头
	fmt.Println(strings.HasSuffix(c, "广西")) // 判断是否为结尾

	// 裁剪字符串，去重
	var s1 = "abcadefg"
	fmt.Println(strings.Trim(s1, "ag"))       // trim两端移除，从字符串两端开始，删除所有属于字符集（参数2）中的字符，直到遇到不属于字符集的字符为止。
	fmt.Println(strings.TrimLeft(s1, "a"))    // 从左端开始
	fmt.Println(strings.TrimRight(s1, "ag"))  // 从右端开始
	fmt.Println(strings.TrimSpace(s1))        // 移除两端的空白字符，中间不管
	fmt.Println(strings.TrimPrefix(s1, "ab")) // 移除前缀，没有就返回原字符串，后缀同理。常用于裁剪格式，如网站、账号、密码等
	fmt.Println(strings.TrimSuffix(s1, "fg")) // 移除后缀

	// 切分字符串，返回类型为string
	var s2 = "www.fblossoms.com"
	fmt.Println(strings.Split(s2, "."))      // 切掉分割符
	fmt.Println(strings.Split(s2, ""))       // 按字符切，不是按字节切，返回类型为string（utf-8）
	fmt.Println(strings.SplitAfter(s2, ".")) // 不切掉分割符
	fmt.Println(strings.SplitAfter(s2, ""))  // 效果与Split一致

	fmt.Println(strings.SplitN(s2, ".", 3)) // n为切完想得到几个字串（0就删除，-1能切就切），元素个数。切分数为 n-1

	fmt.Println(strings.Cut(s2, ".")) // 返回3个值，断点前部分，断电后部分，是否成功
	fmt.Println(strings.Cut(s2, "=")) // 找不到断点就返回，完整字符串，空串，不成功

	// 替换
	fmt.Println(strings.Replace(s2, "com", "edu", 1)) // n为最多替换几次，-1为全部替换（不限次数）

	// 重复
	fmt.Println(strings.Repeat("-", 100)) // Python的print("-" * 100)

	// mao
	fmt.Println(strings.Map(func(r rune) rune { // func(r rune)为匿名函数
		fmt.Printf("%T, %[1]c, %[1]d, %[1]q\n", r)
		if 'a' <= r && r <= 'z' { // 具备函数功能，当前演示为小写转大写
			return r - 32
		}
		return r
	}, s2)) // 相当于for range，按字符返回，进去的值是字符，返回的值也是字符。一一对应
}
