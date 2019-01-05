package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

//func channelCount(chan1 chan int){
//	<-chan1
//	fmt
//}
var desc chan net.Conn
var desc2 chan net.Conn

var conglist []net.Conn

func main() {
	conglist = make([]net.Conn, 0)
	desc = make(chan net.Conn)
	desc2 = make(chan net.Conn)

	srv, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		return
	}
	//go Abosrt()
	go addConn()
	go Abosrt()
	for {
		conn, err := srv.Accept()
		if err != nil {
			return
		}
		//conglist = append(conglist, conn)
		//go
		fmt.Println(len(conglist))
		go Doconnection(conn)
	}

}
func Doconnection(conn net.Conn) {
	desc <- conn
	fmt.Println("connected...")
	//dis-connect..
	//write current info..
	time.Sleep(1e9)

	conn.Write([]byte("hello...\n"))
	conn.Write([]byte(strconv.Itoa(len(conglist)) + " clients..\n"))

	time.Sleep(5e9)

	conn.Close()
	desc2 <- conn

	//send signal...
	//remove..

}

func addConn() {
	for {
		conn1 := <-desc
		//adjust...
		conglist = append(conglist, conn1)
	}

}
func Abosrt() {
	for {
		curCon := <-desc2
		//fmt.Println("del...")
		//del it..
		vL := make([]net.Conn, 0)
		for _, dt := range conglist {
			if dt == curCon {
				//
			} else {
				vL = append(vL, dt)
			}
		}
		conglist = vL
	}
}
