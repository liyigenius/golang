package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func Qsort(list []int) []int {
	if len(list) == 1 {
		return list
	}

	return nil

}

func sumL(List []int) int {
	s1 := 0
	for dt := range List {
		s1 += dt
	}
	return s1
}

func genR(n int) []int {
	l1 := make([]int, 0)
	for i := 0; i < n; i++ {

		l1 = append(l1, rand.Intn(1000))
	}

	return l1
}

func main() {
	l1 := genR(100)
	sort.Ints(l1)
	fmt.Println(l1)

	//l1 := make([]int,0)
	//for i:=0 ; i<=100;i++{
	//	l1 = append(l1, i)
	//}
	//fmt.Println(l1)
	//fmt.Println(sumL(l1))
}
