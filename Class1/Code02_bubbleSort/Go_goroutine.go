package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
	fmt.Println()
}

func gorountinemain() {
	go backGroundTask()

	// 在主goroutine执行其他操作
	for i := 0; i <= 1; i++ {
		fmt.Println("Main Goroutine is running:", i)
		time.Sleep(time.Second) // 等待一秒钟
	}
}

func backGroundTask() {
	//在后台执行一些任务
	for i := 1; i <= 3; i++ {
		fmt.Println("Background Goroutine is running:", i)
		time.Sleep(time.Second) // 等待一秒钟
	}
}

func typewriter(text string, done chan bool) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(100 * time.Millisecond) // 控制打字速度
	}
	fmt.Println()

	done <- true // 发送完成信号
}

func channale() {
	ch := make(chan int)
	v := <-ch
	ch <- v

}
