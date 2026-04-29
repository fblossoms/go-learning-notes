# 命名规范：

- 标识符采用CamelCase驼峰命名法
    - 如果仅在包内可用，就采用小驼峰命名
    - 如果要在包内、包外可见，就采用大驼峰命名
- 简单循环变量可以使用i、j、k、v等
- 条件变量、循环变量可以是单个字母或单个单词，Go倾向于使用单个字母。Go建议使用更短小
- 常量驼峰命名即可
    - 在其他语言中，常量多使用全大写加下划线的命名方式，Go语言没有这个要求
    - 对约定俗成的全大写，例如PI
- 函数/方法的参数、返回值应是单个单词或单个字母
- 函数可以是多个单词命名
- 类型可以是多个单词命名
- 方法由于调用时会绑定类型，所以可以考虑使用单个单词
- 包以小写单个单词命名，包名应该和导入路径的最后一段路径保持一致
- 接口优先采用单个单词命名，一般加er后缀。Go语言推荐尽量定义小接口，接口也可以组合

# 关键字禁止作为标识符：

| 第一列 | 第二列 | 第三列 | 第四列 | 第五列 |
| --- | --- | --- | --- | --- |
| break | default | func | interface | select |
| case | defer | go | map | struct |
| chan | else | goto | package | switch |
| const | fallthrough | if | range | type |
| continue | for | import | return | var |

# 关键字禁止作为标识符：
- 类型 (Types):

| 第一列 | 第二列 | 第三列 | 第四列 |
| --- | --- | --- | --- |
| any | bool | byte | comparable |
| complex64 | complex128 | error | float32 |
| float64 | int | int8 | int16 |
| int32 | int64 | rune | string |
| uint | uint8 | uint16 | uint32 |
| uint64 | uintptr |  |  |

- 常量 (Constants):

| 第一列 | 第二列 | 第三列 |
| --- | --- | --- |
| true | false | iota |

- 零值 (Zero value):

| 第一列 |
| --- |
| nil |

- 函数 (Functions):

| 第一列 | 第二列 | 第三列 | 第四列 |
| --- | --- | --- | --- |
| append | cap | close | complex |
| copy | delete | imag | len |
| make | new | panic | print |
| println | real | recover |  |