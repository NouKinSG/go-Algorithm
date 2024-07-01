package main

import "fmt"

func main() {
	fmt.Println("func start")
	x := 10
	defer func(x int) {
		fmt.Println("in defer: ", x)
	}(x)
	x = 30
	fmt.Println("func end: ", x)
}
