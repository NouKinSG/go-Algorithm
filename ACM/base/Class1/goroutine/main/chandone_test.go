package main

import (
	"fmt"
	"testing"
	"time"
)

func producerTest(out chan<- int, done chan<- struct{}) {
	defer close(out)
	for i := 0; i < 10; i++ { // 假设生成10个数据后停止
		select {
		case out <- i:
			fmt.Printf("producerTest: produced %d \n", i)
		}
		time.Sleep(time.Second) // 模拟生产时间
	}
	// 通知消费者生产完成
	done <- struct{}{}
}

func consumerTest(in <-chan int, done <-chan struct{}) {
	// 无限循环消费
	for {
		select {
		case v, ok := <-in:
			if !ok { // 不ok，说明管道关闭了
				fmt.Println("consumer: channel closed, exiting")
				return
			}
			// ok 可以消费
			fmt.Printf("consumer: received %d\n", v)
			time.Sleep(2 * time.Second) // 模拟处理时间
		case <-done:
			fmt.Println("consumer: producer done, exiting")
			return
		}
	}
}

func TestChan(t *testing.T) {
	ch := make(chan int)        // 创建一个整型通道
	done := make(chan struct{}) // 创建一个通知通道

	go producerTest(ch, done) // 启动生产者 goroutine
	go consumerTest(ch, done) // 启动消费者 goroutine

	// 防止主函数退出
	select {}
}
