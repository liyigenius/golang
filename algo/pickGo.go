package main

import (
	"net"
	"strconv"
)

func a1(a int, b int) int {
	return a + b
}

func main() {
	//fmt.Println(a1(1,2))
	srv, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		return
	}
	count := 0
	userM := make(map[int]net.Conn)
	for {
		conn, err := srv.Accept()
		if err != nil {
			return
		}
		count++
		conn.Write([]byte("hello...you are number " + strconv.Itoa(count)))
		//conn.Close()
		userM[count] = conn
	}
}
