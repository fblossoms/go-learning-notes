# struct{} 空结构体

```go
    var a struct{} => struct{}{}
```

- struct{}{}

```go
    make(chan struct{}, n) n>=0
```

- <- 阻塞了，等信号
  - c <- struct{}{} 等到了不阻塞了
  - close(c) 不阻塞了
  - 不区分，因为往往用这个通道就要阻塞到不阻塞为止
  - 要区分，就要用第二值
