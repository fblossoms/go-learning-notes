package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string // 如果要序列化，也一定要首字母大写，因为要导出属性
	Age  int    `json:"a,omitempty"`
}

func main() {
	p1 := Person{"tom", 20}
	fmt.Printf("%+v\n", p1) // 本质上也是一种序列化，p1打印给你看，依然是字符串，但不是Json
	b, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v, %s\n", b, string(b)) // Json规则的序列化
	fmt.Println("====================================================================")

	var data string = `{"Name":"tom","A":20}` // 使用了标签tag对应的字段需要改正，改正为和标签一致，不区分大小写
	// 拿到的数据，反序列化要转成什么类型，程序员要有数。反序列化有可能是失败的
	// var i int		// object反序列化成int会报错
	var i Person
	err = json.Unmarshal([]byte(data), &i) // 注入scan、Unmarshal时要使用地址
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T %+[1]v\n", i) // i和p1类型相同，实例不同
	fmt.Println(p1 == i)         // 内容相同，Go实现了结构体内容的比较，因为类型相同才能比较内容
	fmt.Println("====================================================================")

	var t interface{}
	err = json.Unmarshal([]byte(data), &t) // var t any = 任意值类型
	if err != nil {
		panic(err)
	}
	fmt.Println(t) // 由于是key-value对链表，顺序可能不一样
	// 所以需要有具体的类型承接转化，而不是使用any类型（interface{}），而且会损坏原有结构和功能
}
