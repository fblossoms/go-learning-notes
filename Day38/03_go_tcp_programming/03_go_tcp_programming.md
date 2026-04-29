# Go TCP编程
- net包，底层已经被修改为nonblocking，使用方式如同阻塞函数一样。由Netpoller 接管，M就去关联的P上获得一个新的G