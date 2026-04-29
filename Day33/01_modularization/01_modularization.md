# 模块化
- 用任何语言开发，如果软件规模扩大，会编写大量的函数、结构体、接口等代码
- 这些代码不可能写在同一个文件中，这就会产生大量的文件
- 如果这些代码杂乱无章，就会造成
  - 命名冲突
  - 重复定义
  - 难以检索
  - 无法引用
  - 共享不变

# 包
- 包由多个文件和目录组成，看见目录就相当于看见了包
- 使用package <包名>来定义包名，每个Go文件必须写明隶属于哪个包的
- 包名一般都采用小写，符合标识符要求
- 当前目录名和package <包名>中的包名不需要一致，但最好保持一致
- 同级文件归属一个包，即每个包目录的当前目录中，只能统一使用同一个package的包名，否则编译出错（测试包例外）
- main函数要在main包内
- 目录里可以新建子目录，相当于包里由子包

# Go Modules
- Go Modules是从Go 1.11版本引入，到1.13版本已经成熟
- 优势
  - 不受GOPATH限制，项目源代码可以放任意目录
  - 统一管理，自动下载依赖，自动清除依赖，且可以控制使用版本（cache目录为GOMODCACHE即GOPATH/pkg/mod）
  - 包管理依赖（go.mod文件）
- 不允许使用相对引入

- GO111MODULE配置
  - 控制Go Module模式是否开启，有off、on、auto（默认）三个值，auto是默认值
    - GO111MODULE=on，支持模块，Go会忽略GOPATH和vendor目录，只根据go.mod下载依赖，在$GOPATH/pkg/mod目录搜索依赖包
      - Go 1.13后默认开启
      - 可以不配置，默认直接开启
    - GO111MODULE=off，不支持模块，Go会从GOPATH和vendor目录寻找包
    - GO111MODULE=auto，在$GOPATH/src外面构建项目且有go.mod文件时，开启模块支持，否则使用GOPATH和vendor机制
  - GOPROXY环境变量可以指定包下载镜像（镜像地址有时会变化，需参考官方文档）
    - https://goproxy.cn,direct
    - https://mirrors.aliyun.com/goproxy/
    - https://mirrors.cloud.tencent.com/go/
    - https://repo.huaweicloud.com/repository/goproxy/

# 包管理
- Go module机制
  - go.mod文件（通过命令创建）

```txt
            go mod init 模块名称
```

    - 告诉第三方包的依赖关系
- import

```txt
        import "ModuleName/目录名"
```

  - ModuleName
    - 项目名称 test（由 go mod init test创建）

```txt
                import "test/crlc"

```

    - 域名/项目名称 gxmu.edu.cn/tools（由 go mod init gxmu.edu.cn/tools创建）

```txt
                import "gxmu.edu.cn/tools/crlc"

```

  - 方式
    - 1、绝对导入
      - 如果是外部的第三方，下载get到GOPATH/pkg/mod

```txt
                    import "fmt"
                    import "go-learning-notes/260319/crlc"
                    import "github.com/vmihailenco/msgpack/v5"
```

    - 2、别名导入

```txt
                    import m "go-learning-notes/260319/crlc"
```

        - m.Add(4, 5) （别名调用）
    - 3、相对导入
      - （已废弃）
    - 4、点导入
      - 把.对应的包的所有导入全局标识符加入到当前名词空间。不推荐，容易造成名词冲突
        - import . "gxmu.edu.cn/tools/crlr"
    - 5、匿名导入
      - 只执行、先执行该依赖的init函数（如果有），不执行该依赖的其他函数，用于调驱动（驱动在init里调用，如创建连接等）

```txt
                    import _ "go-learning-notes/260319/crlc"

```

  - init函数
    - 无参、无返回值，在main函数前执行
    - 同一个文件里可以写多个init函数，顺序不可预期，但是没必要多写，只写一个就多够了

```txt
            import导入某个包的时候，如果有定义init函数就执行，不同包的init函数执行顺序由导入顺序决定

```

    - 匿名导入
      - 没法通过_来访问包内资源，但是会加载包，也会执行init函数
    - 绝对导入、命名导入
      - 也是导入，也会加载包，也会执行init函数，可以通过包名、别名来使用包内资源
