# 字符

```go
type rune = int32 // type定义类型，类型名rune，rune就是int32的别称别名
type byte = uint8 // 类型名byte，byte就是uint8的别称

## Go定义
- 字面量 'a'或'\x61'（十六进制转义） // 1个字符，4个字节，内容为 Unicode码点
- 确定存放的是ASCII里的内容，建议使用byte因为省空间。超过1字节，就只能使用rune