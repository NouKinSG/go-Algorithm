package main

import (
	"Algorithm/ACM/base/Class1/CSP/barrier/message"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var closeChan chan struct{}
var sigs chan os.Signal

func init() {
	sigs = make(chan os.Signal)
	closeChan = make(chan struct{})
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 启动一个 goroutine 监听信号
	go func() {
		select {
		case sig := <-sigs:
			fmt.Println("Received signal:", sig)
		case <-time.After(10 * time.Second): // 在 10 秒后自动关闭通道
			fmt.Println("No signal received, closing after timeout")
		}
		close(closeChan)
		message.ConsumerInst().Exit()
		message.ProducerInst().Exit()
	}()
}

func main() {

	fmt.Println("Main Process begin!")
	<-closeChan
	fmt.Println("main closeChan after!")
	message.ConsumerInst().Join()
	message.ProducerInst().Join()
	fmt.Println("Main Process exit!")
}
