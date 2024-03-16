package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	data string
}

var instance *singleton
var once sync.Once

func getInstance() *singleton {
	once.Do(func() {
		instance = &singleton{data: "Hello, World!"}
	})
	return instance
}

func main() {
	s1 := getInstance()
	s2 := getInstance()

	fmt.Println(s1.data) // 输出: Hello, World!
	fmt.Println(s2.data) // 输出: Hello, World!

	s1.data = "Modified"

	fmt.Println(s1.data) // 输出: Modified
	fmt.Println(s2.data) // 输出: Modified
}
