package main

import "fmt"

func funcFacotry(ftype int) func(a int, b int) int {
	if ftype == 1 {
		return func(a int, b int) int {
			return a + b
		}
	}
	if ftype == 2 {
		return func(a int, b int) int {
			return a * b
		}
	}
	return nil

}

func main() {

	v1 := funcFacotry(1)
	v2 := funcFacotry(2)

	r1 := v1(3, 4)
	r2 := v2(3, 4)
	fmt.Println(r1)
	fmt.Println(r2)

}
