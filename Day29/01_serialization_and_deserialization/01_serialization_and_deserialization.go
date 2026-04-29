package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var data = []any{
		111, // strconv.itoa()
		222,
		333,
		4.44,
		true,
		false,
		[3]int{97, 98, 99}, // 为了解析成数组，会把[和]也算进[]byte类型	// 转成JavaScript的Array
		// []uint8 [91（[的ASCII值） 57 55 44 57 56 44 57 57 93（]的ASCII值）]（这里的中括号是为了表达为byte切片的[]）, [97,98,99]
		// 由Go的json包规定
		// 反序列化后类型变成interface {}（[]any），类型丢失，这是一种兼容效果，因为Go的数据类型种类多余JSON
		map[string]int{"abc": 123, "def": 456},
		// []uint8 [123 34（"的ASCII值） 97 98 99 34 58 49 50 51 44 34 100 101 102 34 58 52 53 54 125]
		map[int]int{1: 123, 2: 234}, // 转成JavaScript的Object
		// []uint8 [123 34（"的ASCII值，int的key也会解析成字符串） 49 34 58 49 50 51 44 34 50 34 58 50 51 52 125]
	}

	var targets [][]byte // 想在切片里面装切片
	// 序列化
	for i, v := range data {
		b, err := json.Marshal(v) // []byte。b为字符对应的ASCII码
		if err != nil {
			continue
		}
		fmt.Printf("%d: %T %[2]v ==> %T %[3]v, %s %[4]q\n", i, v, b, string(b))
		targets = append(targets, b)
	}
	fmt.Println(targets)
	fmt.Println("===========================================================================")

	// 反序列化
	for i, v := range targets {
		var t any
		err := json.Unmarshal(v, &t)
		if err != nil {
			fmt.Println(i, err)
			continue
		}
		fmt.Printf("%d: %T %[2]v ==> %T %[3]v\n", i, v, t)
	}
}
