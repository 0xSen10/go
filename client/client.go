package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("client err:", err)
		return
	}
	fmt.Println("客户端连接成功")
	defer conn.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("终端输入失败 ： err ", err)
		}
		str = strings.Trim(str, " \r\n")
		if str == "exit" {
			fmt.Println("客户端退出")
			break
		}

		n, err := conn.Write([]byte(str + "\n"))
		if err != nil {
			fmt.Println("连接失败， err :", err)
		}
		fmt.Printf("终端数据通过客户端发送成功，一共发送了%d字节的数据，\n", n)
	}
}
