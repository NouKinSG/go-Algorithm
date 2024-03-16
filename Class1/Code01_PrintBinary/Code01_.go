package main

import (
	"fmt"
)

func main() {
	// Your code here
	printBinary(3)
	Constant()
}

func printBinary(num int) {
	for i := 31; i >= 0; i-- {
		if num&(1<<uint(i)) == 0 {
			fmt.Print("0")
		} else {
			fmt.Print("1")
		}
	}
	fmt.Println()
}
