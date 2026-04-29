package main

import (
	"encoding/json"
	"fmt"

	"github.com/vmihailenco/msgpack/v5"
)

type Person struct {
	Name string `json:"n" msgpack:"nn"` // msgpack标签
	Age  int    `json:"a"`
}

func main() {
	var data = []Person{
		{"tom", 4}, {"amy", 5},
	}

	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T %[1]v, %s, len=%d\n", b, string(b), len(b)) // Json规则的序列化
	// json把value转换成数字的字符对应ASCII码，但是由于没有引号，所以最后反序列化后会当成数值理解

	b, err = msgpack.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T %[1]v, %s, len=%d\n", b, string(b), len(b))
	// 二进制序列化把value转换成数字对应ASCII码，就直接是数字。有些数字起到标注的作用

	var i []Person
	err = msgpack.Unmarshal(b, &i)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T %+[1]v, len=%d\n", i, len(i))
}
