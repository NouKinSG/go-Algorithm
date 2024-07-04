package Chnanel

import (
	"fmt"
	"sync"
	"testing"
)

//func TestChannel2(t *testing.T) {
//	ch := make(chan int)
//
//	// 1.往 channel 中发送消息
//	ch <- 2
//
//	<-ch              // 取出消息，直接扔掉
//	value := <-ch     // 从 ch 中读取数据，并保存到 value 变量中
//	value, ok := <-ch // 从ch 读取一个值，判断是否读取成功，如果成功则保存到value变量中
//
//	// // 通过 for range 语法来获取值，直到ch 关闭时退出循环
//	for v := range ch {
//
//	}
//
//}

// sender 不停的向 chan 里面发送数据
func sender(ch chan string) {
	ch <- "hello"
	ch <- "this"
	ch <- "is"
	ch <- "alice"
	// 发送通话结束
	ch <- "EOF"
	close(ch)
}

// recover 循环读取 chan 里面的数据，直到 channel 关闭
func receiver(ch chan string, down chan struct{}) {
	defer func() {
		down <- struct{}{}
	}()

	for v := range ch {
		// 处理通话结束
		if v == "EOF" {
			return
		}
		fmt.Println(v)
	}
}

func TestChannel(t *testing.T) {
	ch := make(chan string) // 创建一个无缓冲通道

	down := make(chan struct{})
	go sender(ch)         // 启动 sender goroutine
	go receiver(ch, down) // 启动 receiver goroutine

	<-down
}

func f3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("In f3")
}

func f2(wg *sync.WaitGroup) {
	defer wg.Done()
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go f3(&wg2)
	fmt.Println("In f2")
	wg2.Wait()

}

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go f2(&wg2)
	fmt.Println("In f1")
	wg2.Wait()

}

func TestMainifact(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go f1(&wg)
	fmt.Println("In main")
	wg.Wait()
}
