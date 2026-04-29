import socket
import threading
import time
import random

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
# application/json

# 模板技术
def get_response():
    with open('index.html', mode="rb") as template:
        data:bytes = template.read()
        body:bytes = data.replace(b"{{}}", str(random.randint(5, 20)).encode())
        return response_header.format(len(body)).replace("\r", "\r\n").encode() + body

# 每个线程处理一个连接，connection per thread
def fn2(conn: socket.socket):  # 把conn看作socket.socket类型，方便做类型检查
    try:
        data = conn.recv(4096)
        if not data: # 数据等效
            print(conn.getpeername(), "byebye")
            return  # 退出当前函数
        print(type(data), data)
        conn.send(get_response())
    except Exception as e: # e = 错误对象
        print(e)
    finally:    # 不管try块中的代码有没有异常错误发生，最终都一定会执行finally
        conn.close()

def fn1(server):
    count = 1
    while True: # 每次有新的客户端接入就启用一个线程调用fn2函数
        conn, raddr = server.accept()
        print('=' * 50)

        threading.Thread(target=fn2, args=(conn,), name=f"r{count}", daemon=True).start()
        # args：收到conn的参数，根据target传到并作为fn2的参数
        count += 1

if __name__ == '__main__':
    # get_response() # byte

    server = socket.socket()
    server.bind(('0.0.0.0', 9999))
    server.listen(1024)

    threading.Thread(target=fn1, name="ac", args=(server, )).start()

    while True:
        time.sleep(10)
        print([t.name for t in threading.enumerate()])