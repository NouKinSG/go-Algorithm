package main

import (
	"Algorithm/ACM/base/Class1/CSP/barrier/message"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var closeChan chan struct{}
var sigs chan os.Signal

func init() {
	//类似于auto
	sigs = make(chan os.Signal)
	//具体类型初始化
	closeChan = make(chan struct{})
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//可以理解为C++ 的匿名函数，或者js的匿名函数，此处通过go原语启动一个协程并行执行
	go func() {
		sig := <-sigs
		fmt.Println("receive signal is ", sig)
		close(closeChan)
		message.ConsumerInst().Exit()
		message.ProducerInst().Exit()
	}()
}

func main() {

	fmt.Println("Main Process begin!")
	<-closeChan
	message.ConsumerInst().Join()
	message.ProducerInst().Join()
	fmt.Println("Main Process exit!")
}
