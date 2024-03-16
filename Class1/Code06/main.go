package main

import (
	"fmt"
	"sync"
)

var flagChan = make(chan struct{})
var wg sync.WaitGroup

func main() {
	// Your code here
	wg.Add(2)
	go work1()
	go work2()
	wg.Wait()
}

// 打印偶数
func work1() {
	for i := 0; i < 100; i++ {
		flagChan <- struct{}{}
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	wg.Done()
}

// 打印奇数
func work2() {
	for i := 1; i <= 100; i++ {
		<-flagChan
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
	wg.Done()
}
