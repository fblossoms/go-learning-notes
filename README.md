# Go 学习笔记与代码示例

**中文** | [English](README.en.md)

内容从 Go 基础语法开始，逐步覆盖数据结构、函数、结构体、接口、错误处理、模块化、日志、序列化、并发、网络编程、数据库访问以及常见 Go 生态库。

这个仓库不是一个单体业务项目，而是一个面向复习、检索和长期积累的学习型代码库。

## 技术栈

- **语言**：Go
- **核心语法**：变量、常量、类型、控制流、函数、指针、结构体、接口、错误处理
- **数据结构**：数组、切片、子切片、哈希表、字符串、线性数据结构
- **并发编程**：goroutine、channel、timer、GMP 模型、通道多路复用
- **网络编程**：TCP 编程、Web Server 基础、HTTP 请求与响应模型
- **数据库**：MySQL、MongoDB
- **数据库工具**：`database/sql`标准库、SQL Builder、GORM
- **日志**：标准库 `log`、`zerolog`、文件日志
- **序列化**：结构体序列化、JSON、二进制序列化、struct tag
- **辅助示例**：Python 并发、socket、简易 Web Server 实验

## 知识地图

### Go 基础

- 注释、标识符、变量、常量、`iota`
- 整型、浮点型、字符串、字符、格式化输出
- 分支语句、循环语句、随机数、输入处理
- 类型转换、作用域、递归、函数嵌套、闭包

### 数据结构与内存模型

- 指针与指针基础
- 数组、切片、子切片、`append`、扩容策略、`copy`
- 线性数据结构、哈希表
- 引用类型、浅拷贝与深拷贝

### 函数、结构体与接口

- 函数定义、函数类型、形参
- `defer` 与执行顺序
- 结构体、结构体指针、匿名结构体
- 构造函数、嵌入结构体、方法接收器
- 接口、类型断言、格式化接口
- 错误处理与基础测试示例

### 模块化、日志与序列化

- Go module 与本地包组织
- 标准日志与 `zerolog`
- 错误日志记录、文件处理
- 结构体序列化、JSON、二进制序列化

### 并发与网络编程

- 并发与并行的基本概念
- goroutine 与 GMP 模型
- TCP 编程
- Web Server 基础示例
- channel、timer、`struct{}` 通道、通道多路复用、通道并发
- Python 并发和 socket 示例作为辅助对照材料

### 数据库与持久化

- 数据库基础概念与 SQL
- MySQL 访问与 `database/sql`
- SQL Builder
- ORM 思想与 GORM
- MongoDB CRUD 示例
- `context` 在数据库和网络类程序中的使用

## 学习路径

1. `Day01` 到 `Day09`：建立 Go 基础语法、基础类型、分支、循环和输入输出概念。
2. `Day10` 到 `Day16`：进入数组、切片、字符串、map、类型转换、哈希表等数据结构主题。
3. `Day17` 到 `Day26`：学习函数、作用域、递归、闭包、`defer`、结构体、接口、拷贝、错误处理与测试。
4. `Day27` 到 `Day34`：整理面向对象风格、序列化、模块化、日志与文件处理。
5. `Day35` 到 `Day40`：集中学习并发、goroutine、channel、TCP 和 Web Server。
6. `Day41` 到 `Day44`：进入数据库相关主题，包括 SQL、MySQL、SQL Builder、GORM、MongoDB 和 `context`。

## 目录结构

每个 `DayXX` 目录下会继续按主题拆分子目录。多数主题目录包含：

- 一个 `.go` 示例文件
- 一个同主题 `.md` 笔记文件

示例：

```txt
Day42/
  01_database_development/
    01_database_development.go
    01_database_development.md
  02_sql_builder/
    02_sql_builder.go
    02_sql_builder.md
```

这种结构可以让每个 `package main` 示例独立编译，避免多个示例放在同一目录时出现 `main redeclared` 报错。

## 运行方式

运行单个主题示例：

```bash
go run ./Day42/02_sql_builder
```

检查整个仓库的 Go 代码：

```bash
go test ./...
```

数据库相关示例需要本地 MySQL 或 MongoDB 服务，以及与代码中连接字符串匹配的数据库配置。

## Skills

- 使用 Go 编写和组织小型可运行示例
- 理解值类型、引用类型和基础内存行为
- 使用数组、切片、map、结构体和接口建模
- 掌握函数、闭包、`defer`、方法接收器和错误处理
- 使用 Go 标准库完成 I/O、时间、日志、SQL、网络等任务
- 使用 goroutine 和 channel 处理并发问题
- 使用 MySQL、MongoDB、GORM、SQL Builder 完成基础持久化操作


