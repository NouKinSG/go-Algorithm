package main

import (
	"fmt"
	"time"
)

// 生产者
func producer(out chan<- int) {
	for i := 0; ; i++ {
		out <- i
		fmt.Printf("producer: produced %d\n", i)
		time.Sleep(time.Second) // 模拟生成时间
	}
}

func consumer(in <-chan int) {
	for v := range in {
		fmt.Printf("consumer: received %d\n", v)
		time.Sleep(2 * time.Second)

	}
}

func main() {
	ch := make(chan int) // 创建一个新型通道

	go producer(ch) // 启动生产者 goroutine
	go consumer(ch) // 启动消费者 goroutine

	// 防止主函数退出
	select {}
}
