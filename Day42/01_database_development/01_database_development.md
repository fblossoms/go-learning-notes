# 数据库开发及驱动
- 驱动包
  - TCP连接及管理
    - 调几个函数就建立好了，无需关心如何维护和管理
  - 协议
    - 如何发请求，接收时如何解析数据
  - 不同语言
- 提供者
  - 官方（优先选择）
  - 第三方
- 共性 抽象
  - 建立连接
  - 断开连接
  - 返回一个数据集，可以遍历

  - Go在你不知道你是什么数据库的驱动的情况下，已经把能操作一个数据库的所有结构体等类型的结构都定义了
    - 包名为database/sql，但不是真正的实现，但是最后一定要落实在具体数据库的驱动包上
    - 驱动包的init中注册到sql.Drivers:map
    - sql.Open("mysql", dsn) --调用-> Driver["mysql"].OpenConnector(dsn)
    - 因此开发时使用sql就行了，不过要把驱动注册一下
  - MySQL 驱动
    - https://github.com/go-sql-driver/mysql  支持 database/sql，推荐
      - 安装指令go get -u github.com/go-sql-driver/mysql
    - https://github.com/ziutek/mymysql       支持 database/sql，支持自定义接口
    - https://github.com/Philio/GoMySQL       不支持 database/sql，支持自定义接口

# 注入攻击

```txt
    s := "123' or '1'='1"
```

- "WHERE username = 'wayne' AND password = '%s'"
- "WHERE username = 'wayne' AND password = '123' OR '1'='1'"
- "WHERE username = ? AND password = ? "

- 参数化查询
  - 最重要的作用就是避免注入攻击
  - 预编译
    - 语句编译并缓存到数据库端，同样的SQL语句不需要再次编译
    - SQL语句都需要在服务端进行 词法分析、语法分析、AST抽象语法树、plan执行计划、优化

- 数据库最耗时的是慢查询、慢操作
