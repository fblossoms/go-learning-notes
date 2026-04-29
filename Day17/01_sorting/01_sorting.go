package main

import (
	"fmt"
	"sort"
)

func main() {
	// 升序
	a := []int{100, 233, 12, 346, -22, 67} // 在物理地址上是有先后顺序的，所以当前是有序的
	fmt.Printf("%T, %[1]v, len=%d, cpa=%d\n", a, len(a), cap(a))
	sort.Ints(a)                                                 // 调用sort包，使用Ints，因为当前切片a为int类型。不能赋值，因为当前函数无返回值
	fmt.Printf("%T, %[1]v, len=%d, cpa=%d\n", a, len(a), cap(a)) // 就地排序，直接改了a的切片
	//sort.Sort(sort.IntSlice(a)) // IntSlice在sort包（外包）中定义（的新类型。为了扩展新方法 ），所以要使用sort,指明来源

	b := []string{"", "ABC", "s", "123", "Abc", "abc", "bac"}
	fmt.Printf("%T, %[1]v, len=%d, cpa=%d\n", b, len(b), cap(b))
	sort.Strings(b)
	fmt.Printf("%T, %[1]v, len=%d, cpa=%d\n", b, len(b), cap(b))

	// 降序
	sort.Sort(sort.Reverse(sort.IntSlice(a)))                    // 输入IntSlice后自动跳出来的{}只能输入字面常量，所以要自行换成变量
	fmt.Printf("%T, %[1]v, len=%d, cpa=%d\n", a, len(a), cap(a)) // 也是就地排序

	sort.Sort(sort.Reverse(sort.StringSlice(b)))
	fmt.Printf("%T, %[1]v, len=%d, cpa=%d\n", b, len(b), cap(b))

	// 二分查找：前提是必须先排好升序再进行二分查找
	// 默认内置了排升序算法。若已排好升序，则宣先进行比较，确认后就进行二分查找。确认后不会再进行排序操作
	c := []int{55, 66, 11, -22, -11, 0}
	sort.Ints(c)
	index := sort.SearchInts(c, 5)
	fmt.Printf("%T, %[1]v, len=%d, cpa=%d\n", c, len(c), cap(c))
	fmt.Println(index)

}
