package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name  string
	score int
}

// 扩展方法
type StudentSlice []Student

func (a StudentSlice) Len() int      { return len(a) }
func (a StudentSlice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// 自定义类型难以直接比较大小，需要先定义比较规则。这里以比较成绩为规则
func (a StudentSlice) Less(i, j int) bool { return a[i].score < a[j].score }

func main() {
	// 结构体struct排序
	s1 := Student{"feifei", 100}
	s2 := Student{"Xixi", 70}
	s3 := Student{"Jiejie", 80}
	students := []Student{s1, s2, s3}

	fmt.Println(students)
	// 升序
	sort.Sort(StudentSlice(students)) // 就地排序（升序）
	fmt.Println(students)
	// 降序
	sort.Sort(sort.Reverse(StudentSlice(students))) // Reverse原理是将i和j调换
	fmt.Println(students)

	// 灵活写法（无需自定义类型）
	sort.Slice(students, func(i, j int) bool {
		return students[i].score < students[j].score // 按成绩降序
		// return students[i].name < students[j].name // 按名字降序
	})

	// map排序
	// map排序的思路是按照key排序或按照value排序，但是map的key和value是结合在一起的
	m := map[string]int{"Feifei": 100, "Xixi": 90, "Jiejie": 95}
	fmt.Println(m) // 哈希表本身就是无序的

	// 基于key排序：把key装到切片里进行排序
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	fmt.Println(keys) // 哈希表本身就是无序的
	// 升序
	sort.Strings(keys)
	fmt.Println(keys)
	// 取到value
	for _, k := range keys {
		fmt.Println(k, m[k])
	}

	// 基于value排序（明确key）：把key和value一起装进线性容器里进行排序
	type Entry struct {
		K string
		V int
	}
	var entires []Entry
	for k, v := range m {
		entires = append(entires, Entry{k, v})
	}
	fmt.Println(entires)
	// 升序
	sort.Slice(entires, func(i, j int) bool {
		return entires[i].V < entires[j].V
	})
	fmt.Println(entires)
}
