import socket

server = socket.socket()  # socket（包）.socket（类），socket包下的socket类实例化获得一个socket对象
# socket对象也是文件，有文件描述符（0：标准输入，1：标准输出，2：标准错误输出）
# 默认TCP、IPv4

# 端口是一个双字节无符号整数，[0, 65535]，共65536种状态，建议1024下别随便使用
server.bind(("0.0.0.0", 9999))  # socket套接字格式 IP:port。0.0.0.0表示所有本机可用IP地址（包括回环地址）

# 监听，三次握手后建立的TCP连接会放在监听队列里
server.listen(1024) # 非阻塞函数
print(server)
print('-' * 50)

# 如果有客户端连接到来，你要接进来
# accept盯着监听队列里面的已完成握手的连接，只要有建立好的连接
# accept调用一次就从队列里拿一个连接对象
# 为这个连接分配一个可以操作的对象socket
# 这个连接对象的另一端就是客户端
# 这个不妨碍底层使用TCP协议
conn, raddr = server.accept()   # 接客了，门童，迎接客户端，阻塞函数，谁阻塞？当前线程阻塞。返回值为元组类型 (socket对象, 地址信息)
print(conn)
print(raddr)
print(conn.getpeername(), raddr)    # 拿到对端地址
print(conn.getsockname())   # 拿到本端地址
print('=' * 50)

for i in range(4):  # 收发三次echo server
    data = conn.recv(4096) # recv为接收，也是read读取，默认阻塞，等待数据
    print(type(data))
    conn.send(b"hello! feifei!") # 发送

conn.close()
server.close()