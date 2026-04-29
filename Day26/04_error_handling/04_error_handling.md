# 错误处理
- 状态码
  - goroutine 1 [running]:
  - 0为正常退出，非0均为异常

- 自定义error
  - 系统报出的错误不一定是最终的错误，可能只是系统的一种表达
  - Go不分异常和错误，都归于错误
    - 避免自己写出bug
    - 考虑潜在的错误
  - Go的错误处理
    - 函数调用，返回值可以是返回多值，一般最后一个值可以是error接口类型的值
      - 如果函数调用产生错误，则这个值是一个error接口类型的错误
      - 如果函数调用成功，则该值是nil
    - 检查函数返回值的错误是否是nil，如果不是nil，有必要进行错误处理

  - 其他语言（Java等）
    - try {
      - 多条语句（有可能抛出异常或错误）
    - } catch {
      - 抓取指定的错误
    - } finally {
      - 写不管有没有错误，最终都要指向的代码（资源的释放）

```txt
            }
```

    - 错误或异常是用类、子类实现的
  - Go语言
    - 产生一个错误
      - 1、编写代码中使用panic，手动主动地抛出错误
      - 2、不是程序员自己写地，内部产生一个错误 panic
    - 全局error接口
      - 一个全局方法 error() string，返回的字符串就是错误的描述原因
      - 如果要自定义一个错误，那么请实现error接口

```go
                    func foo() (int, error（错误接口类型）) {}
```

      - 自定义非接口类型时使用e或err

```txt
                    v, err := foo()
                    if err != nil {
```

          - // 对错误进行处理
        - } else {}

  - 源码

```go
            func New(text string) error {
            	return &errorString{text}   // errorString类型的实例的指针，实现了Error方法，实现了error接口，指针既是又是error接口类型的
            }

            type errorString struct {
```

      - s string

```go
            }

            func (e *errorString) Error() string {
            	return e.s
            }

```

    - 即 err := New("原因及描述") => *errorString类型的，又是error接口类型 => 又是err.Error()方法

- panic宕机
  - 具体问题具体分析
  - recover可用处理掉panic，也可以不管

- recover
  - 内建函数
  - 写在defer后面的函数中

  - runtime.Error接口
    - 和error接口的关系
      - 内嵌了error类型，相当于是子类型，实现了runtime.Error接口，就一定实现了error接口
    - runtime.errorString关系

    - 源码1

```go
                type Error interface {  // runtime包
```

        - error   // 内嵌，实现Error() string
        - RuntimeError

```txt
                }
```

    - 源码2
      - type errorString string // errorString 实现error接口，同时也实现了runtime.Error接口

```go
                func (e errorString) RuntimeError() {}
                func (e errorString) Error() string {
                    return "runtime error: " + string(e)
                }
```

    - 示例

```txt
                var x errorString = "abc"
```

      - x 既是runtime.errorString，又是error接口类型的，又是runtime.Error类型的

- 面向对象的实现就是依赖结构体的接口实现
  - 开发项目中，首先约定接口，操作、方法定义好
  - 然后去定义结构体，实现某些接口的方法
  - 用 接口类型，而不是 结构体类型。所以调用的是该接口类型定义的方法