package main

import (
	"fmt"
)

func main() {
	fmt.Println("start main")
	fn2()
	fmt.Println("end main")
}

func fn2() {
	fmt.Println("start fn22222")
	fn()
	fmt.Println("end fn22222")
}

func fn() {
	fmt.Println("start fn")
	panic("pannic in fn")
	fmt.Println("end fn")
}

// 会引发panic的sum
func sum(x, y *int) int {
	return *x + *y
}
