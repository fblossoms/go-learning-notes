# 上下文Context（标准库，很有用，必须会）
- 主要用在父子协程间的控制，在上下游协程中传递信号或数据
- 是一个接口的定义
- type emptyCtx int
  - 实现了Context接口的4个方法，虽然偷懒了，都是空实现
- todo = new(emptyCtx) 通过TODO()拿到 我不知道用不用上下文，随便指定一个
- background = new(emptyCtx) 通过 BackGround()，根的上下文一般采用BackGround作为根上下文
- 既是emptyCtx类型的指针，又是Context接口类型的指针

- 自定义的上下文对象都要基于BackGround，BackGround子子孙孙上下文
  - 继承，祖先 实现了 Context的4个接口方法，子孙就可以不实现，直接继承
  - 子孙上下文 覆盖 这4个方法中某些方法

- valueCtx值上下文

```go
        type valueCtx struct {
```

    - Context // 匿名属性，嵌套，父子结构体，父Context 接口类型，属性名字Context，继承4个方法的
    - // emptyCtx
    - key any
    - val any

```txt
        }
```

  - 覆盖了Value方法
  - WithValue(parent, key, val) Context // 构建基于parent上下文对象的子上下文实例

```txt
            return &context.valueCtx{Context:Backgroud(), key:"name", val:"tom"}

```

- cancelCtx
  - Done() 返回只读的同步通道`<- chan struct{}`，该通道会在上下文被取消后关闭

```txt
        context.WithCancel(parent)
```

    - WithCancel能构建一个被取消的的上下文对象，同时返回一个无参的取消函数（*cancelCtx和一个函数具有调用cancel的能力）供用户使用
    - cancel函数内部上下文对象的cancel方法即cancelCtx.cancel方法
    - 传播
      - 父上下文调用cancel函数时，会调用其所有子上下文的cancel函数
- timerCtx
  - 计时器


# 空实现源代码

```go
    type emptyCtx struct{} // emptyCtx
    type backgroundCtx struct{
```

  - emptyCtx // emptyCtx类型，说明是嵌套、父子结构体、面向对象继承，属性名emptyCtx，类型emptyCtx，继承了数据和方法

```go
    }
    func (emptyCtx) Deadline() (deadline time.Time, ok bool) {
        return
    }
    func (emptyCtx) Done() <- chan struct{} {
        return nil
    }
    func (emptyCtx) Err() error {
        return nil
    }
    func (emptyCtx) Value(key any) any {
        return nil
    }
    func Background() Context {
        return backgroundCtx{} // 返回实例零值，既是BackgroundCtx类型的，又是Context类型的
    }

```
