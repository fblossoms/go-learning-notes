# 线程池

import socket
import threading
import time
import random
from concurrent.futures import ThreadPoolExecutor

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

def get_response() -> bytes:
    html = response_body.format(random.randint(100, 300)).encode()  # 动态网页技术举例
    response = response_header.format(len(html)).replace('\n', '\r\n').encode() + html
    return response

# 每个线程处理一个连接，connection per thread
def fn2(conn: socket.socket):  # 把conn看作socket.socket类型，方便做类型检查
    try:
        data = conn.recv(4096)
        # 里面的路径或者参数或提交的数据不同对应不同的HTML或者数据 后端路由
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

        # threading.Thread(target=fn2, args=(conn,), name=f"r{count}", daemon=True).start()
        executor.submit(fn2, conn)
        count += 1

if __name__ == '__main__':
    # 线程池
    # 不会一下子创建完2000个，来一个创建一个，会预留空闲
    executor = ThreadPoolExecutor(max_workers=2000) # 最大线程池数2000。懒？是懒的，来一个创建一个

    server = socket.socket()
    server.bind(('0.0.0.0', 9999))
    server.listen(1024)

    # threading.Thread(target=fn1, name="ac", args=(server, )).start()
    # 注册执行
    # 不用写名字，其实名字不重要，名字只是给程序员看的
    executor.submit(fn1, server)

    while True:
        time.sleep(10)
        print([t.name for t in threading.enumerate()])