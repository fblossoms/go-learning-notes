package main

/*
Python 代码

c = "测"
print("Unicode码点10进制表示:", ord(c))  # Unicode
print("Unicode码点16进制表示:", hex(ord(c)))  # 0x6d4b

print("\n----以下是各种编码的码点----")
print("GBK   :", c.encode("gbk"))
print("UTF-8 :", c.encode())  # 默认utf-8编码, b'\xe6\xb5\x8b'
print("UTF-16:", a:=c.encode("utf-16"), [hex(b) for b in a])  # BOM 表明字节序的
print("UTF-32:", a:=c.encode("utf-32"), [hex(b) for b in a])
print("\x6d", "\x4b")
print('~' * 30)
print(int.from_bytes(b'\x6d\x4b', "big"))
print(int.from_bytes(b'\x4b\x6d', "little"))

*/
