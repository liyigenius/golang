package main

import (
	"fmt"
	"math/rand"
	"time"
)

var chan1 chan int

func suck() {

	for {
		v1 := <-chan1
		fmt.Println(v1)
	}

}

func add() {
	for {
		v1 := rand.Intn(100)
		chan1 <- v1
	}
}

func main() {
	chan1 = make(chan int)
	go suck()
	go add()
	go add()

	time.Sleep(1e9)

}
