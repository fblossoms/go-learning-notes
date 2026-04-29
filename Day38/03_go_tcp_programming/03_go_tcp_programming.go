package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	laddr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:9999")
	if err != nil {
		log.Fatal(err) // 直接调os.Exit(1)
	}

	server, err := net.ListenTCP("tcp4", laddr) // 底层非阻塞socket创建，bind、listen
	if err != nil {
		log.Panic(err)
	}
	defer server.Close()

	for {
		conn, err := server.Accept() // 非阻塞的，用法是阻塞的用法
		if err != nil {
			fmt.Println(err)
			// continue
		}
		//defer conn.Close()

		buffer := make([]byte, 4096)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			// break return
		}
		if n == 0 { // 客户端主动断开
			conn.Close()
			continue
		}
		data := buffer[:n] // buffer可能有些空间没用到，就取到n大小
		fmt.Println(data)

		conn.Write(data)
		conn.Close()
	}
}
