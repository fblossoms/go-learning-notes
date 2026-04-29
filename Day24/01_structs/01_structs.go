package main

import "fmt"

// 全局变量
// type user struct { // user为自定义的标识符（结构体名称），struct以下称为“本体”
// 	   fn    func()
// 	   id    int
// 	   name  string
// 	   addr  *int
// 	   score map[int]string
// }

type user struct { // 结构体（这里user），若要在包外使用该结构体，注意首字母要大写（这里是User）
	id    int // 属性名（这里是id），若要在包外使用该字段，注意首字母要大写（这里是Id）
	name  string
	score int
}

type Information struct {
	Id        int
	Stuname   string
	Classname string
	Score     int
}

// 普通函数调用，传入的是Informtion类型的值，取Stuname属性
func getUserName(u Information) string {
	return u.Stuname // 要使用结构体（Information）的实例来取，如果要用实例又得先用传参的方式传入变量并指定类型
}

// 推荐使用
// method：Go提供一种语法，不需要提供参数也能够访问其属性
// 有receiver这种方式的函数称为该recevier类型的方法
// 按照这种语法扩展方法，实例.方法名(实参)调用时，实例和receiver的绑定
func (u Information) getName() string {
	return u.Stuname
	// return getUserName(u) // 可以这样写
	// return u.getUserName() // 不可以，这是普通函数，不是成员
	return u.getName() // 不可以，无限递归了
}

func (u Information) getNameWithPrefix(prefix string) string {
	return fmt.Sprintf("%s == %s", prefix, u.Stuname)
}

func main() {
	// 定义到main里面就是局部变量
	// 定义未来结构体如果要使用，就要有这些数据
	// type user struct { // user为自定义的标识符（结构体名称），struct以下称为“本体”
	// 	   fn    func() // {}里面是数据，一般称为字段field，属性property
	// 	   id    int
	// 	   name  string
	// 	   addr  *int
	// 	   score map[int]string
	// }

	// 1
	var u1 user // 零值，在内存中构建出来user的对应结构，零值大胆用，所有字段都是零值
	fmt.Printf("%T\n", u1)
	fmt.Println(u1)
	fmt.Printf("%v\n", u1)
	fmt.Printf("%+v\n", u1) // 带字段名打印
	fmt.Printf("%#v\n", u1) // 带类型名和字段名打印

	// 2 字面量
	var u2 = user{}
	fmt.Printf("%#v\n", u2)

	u3 := user{20234332916, "feifei", 90} // 采用该方式时，必须给所有字段赋值
	fmt.Printf("%T %+[1]v\n", u3)

	u4 := user{id: 20234332916} // 使用字段给值时可以不给出所有字段，没给出的字段使用零值
	fmt.Printf("%T %+[1]v\n", u4)

	u5 := user{score: 100, name: "feifei", id: 20234332916} // 使用字段给值时可以不按照顺序，输入时会帮自动排好
	fmt.Printf("%T %+[1]v\n", u5)

	u6 := Information{Score: 100, Stuname: "feifei", Id: 20234332916, Classname: "Math"} // 使用字段给值时可以不按照顺序，输入时会帮自动排好
	fmt.Printf("%T %+[1]v\n", u6)
	fmt.Println(u6.Id, u6.Stuname, u6.Classname, u6.Score)

	// 函数返回
	fmt.Println(getUserName(u6))

	// method
	fmt.Println(u6.getName())

	fmt.Println(u6.getNameWithPrefix("***"))

	var user1 Information // 显然，零值可以
	fmt.Println(user1.getName())
	fmt.Println(user1.getNameWithPrefix("###"))

	// 观察函数签名
	fmt.Printf("%T\n", u6.getName)          // 不加括号相当于不取调用函数的结果，而是取函数
	fmt.Printf("%T\n", Information.getName) // 不加括号相当于不取调用函数的结果，而是取函数

}
