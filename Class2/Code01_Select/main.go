package main

import (
	"fmt"
	"time"
)

// Select的出现是为了解决，多个管道的选择问题。多路复用，在多个管道中随机公平的选择一个来执行
/*
	PS: case后面必须进行的是io操作，不能是等值，随机去选择一个io操作
	PS：default 防止select被阻塞住，加入default
*/

// 案例：

func main() {

}

func learnChan() {
	// Your code here

	// 定义 一个 int管道
	intChan := make(chan int, 1)
	go func() {

		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 1)
			fmt.Printf("int通道休眠:%v秒\n", i+1)
		}
		intChan <- 10
	}()

	// 定义一个 string 管道
	stringChan := make(chan string, 1)
	go func() {
		for i := 0; i < 2; i++ {
			time.Sleep(time.Second * 1)
			fmt.Printf("string通道休眠:%v秒\n", i+1)
		}
		stringChan <- "tzjgolang"
	}()
	// fmt.Println(<-stringChan) // 取数据本身就是阻塞的

	select {
	case v := <-intChan:
		fmt.Printf("intChan通道取出的数据是:%v\n", v)
	case v := <-stringChan:
		fmt.Print("stringChan通道取出的数据是:%v\n", v)
	default:
		fmt.Println("default操作防止堵塞")
	}
}

func fibonacci(c, quit chan int) {
	x, y := 1, 0
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
