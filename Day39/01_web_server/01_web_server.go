package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

const response_body = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Web Server-Goroutine</title>
</head>
<body>
    <h2 style="color: blue;">Welcome to GXMU</h2>
    <hr>
    <img src="" alt="no img">
</body>
</html>`

const response_header = `HTTP/1.1 200 OK
Date: Mon, 24 Oct 2024 20.03.23 GMT
Content-Type: text/html; charset=utf-8
Content-Length: %d
Connection: keep-alive
Server: cas.gxmu.edu.cn

%s`

var response = fmt.Sprintf(strings.ReplaceAll(response_header, "\r", "\r\n"), len(response_body), response_body)

func main() {
	fmt.Println(response)
	laddr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:9999")
	if err != nil {
		log.Fatal(err)
		return
	}

	server, err := net.ListenTCP("tcp4", laddr)
	if err != nil {
		log.Panic(err)
		return
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			continue // goto END
		}

		go func(conn net.Conn) {
			defer conn.Close()
			buffer := make([]byte, 4096)
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println(err)
			}
			if n == 0 {
				return
			}
			data := buffer[:n] // 相当于浏览器发来的HTTP请求，根据URL method -> 不同的函数处理返回响应的数据
			// request_first, _, _ := strings.Cut(string(data), "\r\n")
			// fmt.Println(request_first)
			// strings.Split(request_first, " ")

			fmt.Println(data)
			conn.Write([]byte(response))
		}(conn)

	}
}
