import random
import selectors
import socket

response_body = """\
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>测试网页</title>
</head>
<body>
    <h1 style="color: red">欢迎访问 GXMU-{}</h1>
</body>
</html>   
"""

response_header = """\
HTTP/1.1 200 OK
Content-Type: text/html; charset=utf-8
Content-Length: {}
X-Server: 
Connection: keep-alive

"""

def get_response() -> bytes:
    html = response_body.format("IO多路复用，事件驱动").encode()
    response = response_header.format(len(html)).replace('\n', '\r\n').encode() + html
    return response

def fn1(server):
    conn, raddr = server.accept()
    conn.setblocking(False)
    print(conn, raddr, "~~~~~")
    k = st.register(conn, selectors.EVENT_READ, fn2)
    print(k)

def fn2(conn):
    try:
        data = conn.recv(1024)
        print(data, "@@@@") # data浏览器，httprequest 报文，解析
        if not data:
            print("bye", conn.getpeername())
            return
        conn.send(get_response())
        print("+" * 50)
    except Exception as e:
        print(e)
    finally:
        st.unregister(conn) # 关闭前先要反注册
        conn.close()

if __name__ == '__main__':
    # 默认选择器：自动匹配操作系统，选择最优的。Windows只能用select技术
    st = selectors.DefaultSelector()
    # 把关注的IO事件注册给选择器

    server = socket.socket()
    server.setblocking(False)   # 设置为非阻塞IO，关注的IO事件都要注册成non-blocking
    server.bind(('0.0.0.0', 9999))
    server.listen()
    print(server)

    # st.register(server.fileno())
    k = st.register(server, selectors.EVENT_READ, fn1)
    print(type(k), k)

    # st.select() # 选择器监控，默认是阻塞的
    while True:
        for key, mask in st.select():  # 选择器监控，默认是阻塞的
            print(key.fileobj, key.fd, key.events, key.data)
            key.data(key.fileobj)

    # conn, raddr = server.accept() # 可读时再调用
    # print(type(conn), conn)
    #
    # conn.close()
    # server.close()

    print("+" * 50)