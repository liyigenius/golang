package main

import "fmt"

var cache1 map[int]int

func getMin(l []int) int {
	min := 0
	for _, dt := range l {
		if dt < min || min == 0 {
			min = dt
		}
	}
	return min

}

func getCoinN(i int) int {
	if dt, ok := cache1[i]; ok {
		return dt
	}

	if i == 1 || i == 2 || i == 5 {
		return 1
	}
	l1 := make([]int, 0)
	for i1 := 1; i1 < i/2+1; i1++ {
		l1 = append(l1, getCoinN(i1)+getCoinN(i-i1))
	}
	retV := getMin(l1)
	cache1[i] = retV
	return retV

}

func main() {
	cache1 = make(map[int]int)
	for i := 1; i < 101; i++ {
		fmt.Println(i, getCoinN(i))
	}
}
