# fmt.Print*  库函数内部实现
- 结构体默认打印格式   Print，Printf %v，Println
  - 实例 %v       {"tom", 20}
  - 实例指针 %v   &{"tom", 20}
- 结构体打印格式符    %+v 带字段名不带类型
  - 实例 %v       {name:Tom age:20}
  - 实例指针 %v   &{name:Tom age:20}
- 结构体打印格式符    %#v 带字段名带类型
  - 实例 %v       main.Person{name:"Tom", age:20}
  - 实例指针 %v   &main.Person{name:"Tom", age:20}

# 为结构体实现一个方法 String() string，影响 %v和 %+v

# fmt.Println(1, m, &m)的原理
- a []any{1, m, &m}

```txt
    for i, v := a {
```

  - // 进行 v any 断言

```txt
    }
```