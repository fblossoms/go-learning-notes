package main

import (
	"crypto/md5"
	"crypto/sha512"
	"fmt"
)

func main() {
	hash := md5.New() // 。New()创建一个 MD5 计算器对象
	hash.Write([]byte("abc"))
	r := hash.Sum(nil)                              // 不追加就给nil
	fmt.Printf("%T, %[1]v, %[1]s, %d\n", r, len(r)) // 16byte即128bit
	s := fmt.Sprintf("%x", r)
	fmt.Printf("%s, len=%d\n", s, len(s))

	hash2 := sha512.New() // 。New()创建一个 sha512 计算器对象
	hash2.Write([]byte("abc"))
	r2 := hash2.Sum(nil)                              // 不追加就给nil
	fmt.Printf("%T, %[1]v, %[1]s, %d\n", r2, len(r2)) // 16byte即128bit
	s2 := fmt.Sprintf("%x", r2)
	fmt.Printf("%s, len=%d\n", s2, len(s2))

	// 哈希表
	//var m1 map[int]int // 未初始化哈希表，格式为map[key键的数据类型]value值的数据类型，但未赋值nil map不可用
	//fmt.Println(m1 == nil)

	m1 := map[int]int{100: 111, 123: 222} // 初始化map：字面量定义，有底层桶数组（此时B=0，即有一个桶）
	m1[101] = 30
	fmt.Printf("%T, %[1]v\n", m1) // 展示时key是有序的，但其实在内存中是无序的
	m1[123] = 111                 // 覆盖原键值对的value
	fmt.Printf("%T, %[1]v\n", m1)

	m2 := make(map[string]int, 100) // 初始化map：make定义，有底层桶数组，100为期望存储的键值对数量（此时B根据LoadFactor匹配）
	m2["a"] = 32
	fmt.Printf("%T, %[1]v\n", m2)              // 展示时key是有序的，但其实在内存中是无序的
	m2["bcd"] = 96                             // 可以是空船，但是不建议
	m2["a"] = 23                               // 覆盖
	fmt.Printf("%T, %[1]v, %d\n", m2, len(m2)) // 可以取len()，但是不能用cap()
	fmt.Println(m2["c"])                       // 没有key也不会报错，返回零值
	value, ok := m2["d"]                       // 想知道key是否真的不存在
	if ok {
		fmt.Println("key存在")
	} else {
		fmt.Println("key不存在, 零值为", value)
	}

	// 删除
	delete(m2, "a")
	fmt.Printf("%T, %[1]v, %d\n", m2, len(m2)) // 可以取len()，但是不能用cap()
	delete(m2, "c")                            // 如果key不存在，也不会报错，想知道是否有这个key，需要先在delete前自行做判断

	// 遍历，只能for range
	for k, v := range m1 {
		fmt.Printf("%d:%d\t%d\n", k, v, m1[k]) // 可以用key遍历
	}

	// 获取所有key，装到切片里
	keys := make([]int, 0, len(m1))
	for k := range m1 {
		keys = append(keys, k)
	}
	fmt.Println(keys) // 可以看到顺序与前面不一样，印证了在内存中是无序的

	// 获取所有value，装到切片里
	values := make([]int, 0, len(m1))
	for /*k,*/ v := range m1 {
		values = append(values, v)
		//values = append(values, m1[k]) // 也可以通过key取
	}
	fmt.Println(values)

	// 清除：清除所有键值对。原理为直接让链表指针悬空，B不改，count置零
	// 清除大规模数据会引起大范围垃圾回收，降低程序运行使劲按
	fmt.Printf("%T, %[1]v, %d\n", m2, len(m2))
	clear(m2)
	fmt.Printf("%T, %[1]v, %d\n", m2, len(m2))
	s3 := []int{1, 2, 3, 4, 5, 6} // clear可以清除 map 和 slice
	fmt.Printf("%T, %[1]v, %d, %d\n", s3, len(s3), cap(s3))
	clear(s3) // 元素置零
	fmt.Printf("%T, %[1]v, %d, %d\n", s3, len(s3), cap(s3))

}
