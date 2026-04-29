# 循环

Go中只有for循环，没有while循环。for循环风格与C语言相近

- for 初始化执行1次; 循环条件; 每一趟执行完的语句 {}

  - 初始化执行：循环开始之前，一定会做一次，仅一次。该部分可能会引起死循环
  - 循环条件：每一趟执行都要满足该条件；如果该条件测试失败则循环终止。该部分可能会引起死循环
  - 执行完语句：每一趟执行完执行的语句。该部分可能会引起死循环

  - 举例：

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

  1、setup init i := 说明i只能在当前for中用, i现在是0  
  2、看条件, 0 < 5 true进入循环体, 进入循环,执行一趟  
  3、第1趟执行完成后, i++, i是1  
  4、看条件, 1 < 5 true进入循环体, 进入循环, 执行一趟  
  5、第2趟执行完成后, i++, i是2  
  6、看条件, 2 <5 true进入循环体, 进入循环, 执行一趟  
  7、第3趟执行完成后, i++, i是3  
  8、看条件, 3 < 5 true进入循环体, 进入循环, 执行一趟  
  9、第4趟执行完成后, i++,i是4  
  10、看条件, 4 < 5 true进入循环体, 进入循环, 执行一趟  
  11、第5趟执行完成后, i++, i是5  
  12、看条件, 5 < 5 false, 不进入循环体, 循环结束

  - 死循环格式：

```go
for ; ; {}
for {}      // 推荐
for true {}
```

  - 循环可以嵌套，建议1到2层为宜，不要超过三层。总循环次数（时间复杂度）为 i * j *...（i，j，...为每层循环次数）

  - continue和break
    - continue：停止这一趟循环，执行这一趟结束后的for循环部分，执行时需要判断条件是否能够进入循环
    - break：打破最近一层循环，直接执行for后面的语句
    - continue/break和if没有任何关系，只认for

  - 注意关注作用域问题：
    - i为全局变量还是局部变量

- for ... range 高级for

  - range 容器里面有若干元素

    - 线性表：如字符串、数组、切片等

```go
for index, value := range 容器
for _, value := range 容器    // 省略index的写法
for index, _ := range 容器    // 省略value的写法
```

      - 举例：

```go
a := "abcd测试"
for i, v := range a {   // 遍历时，按照字符rune（int32）十六进制ASCII，汉字时使用unicode编码的码点
    fmt.Println(i, v) // v为十六进制对应的十进制的ASCII值（utf-8兼容ASCII）
}
```

      - for-range使用unicode编码
      - len使用utf-8编码
      - 两者存在转换

    - 映射（map，字典）

```go
for key, value := range map
for _, value := range map    // 省略key的写法
for key, _ := range map      // 省略value的写法
```

    - 管道（channel）

```go
for value := range channel
```

# 字符串（拓展）

- 线性表
  - Go语言中，字符串底层数组使用了utf-8编码，每个汉字使用了3字节，英文兼容ASCII每个占用1字节（utf-8兼容ASCII）
  - 文字、标点等一切符号，计算机都不能直接使用，计算机只认识二进制数，所以需要转换为统一标准的二进制数字