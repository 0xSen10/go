package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器的read err ：", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器启动完成,开始监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("启动失败,err=", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端链接：")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("客户端连接失败，err:", err)
			return
		} else {
			fmt.Printf("客户端连接成功，conn:%v  客户端：%v", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
}
