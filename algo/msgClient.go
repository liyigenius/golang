package main

import (
	"fmt"
	"net"
)

//func link(conn net.Conn){
//
//
//	for{
//
//	}
//}

func main() {

	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}
	//go link(conn)

	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		fmt.Printf("Received data: %v", string(buf[:len]))
	}

}
