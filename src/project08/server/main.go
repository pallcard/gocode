package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	
	defer conn.Close()
	
	for {
		buf := make([]byte, 8096)
		fmt.Println("读取客户端发送的数据.....")
		_, err := conn.Read(buf[:4])
		if err != nil {
			fmt.Println("conn.Read err=", err)
			return
		}
		fmt.Println("读到的buf=", buf[:4])
	}
}

func main() {

	fmt.Println("服务器在8889端口监听.....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")

	defer listen.Close()

	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}

	for {
		fmt.Println("等待客户端来链接服务器.....")

		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accpet err=", err)
		}

		go process(conn)

	}

}