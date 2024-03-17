package main

import (
	"fmt"
	"sync"
)

func main() { // Your code here
	fmt.Println("0000")
	once := &sync.Once{}
	var wg sync.WaitGroup // 声明一个WaitGroup变量

	for i := 0; i < 4; i++ {
		wg.Add(1) // 为每个goroutine增加等待计数
		i := i
		go func() {
			defer wg.Done() // goroutine完成时减少等待计数
			once.Do(func() {
				fmt.Printf("first %d\n", i)
			})
			fmt.Println("1111")
		}()
	}

	wg.Wait() // 等待所有goroutine完成

	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	time.Sleep(time.Second) // 模拟一些工作
	// 	wg.Done()
	// }()

	// go func() {
	// 	wg.Wait() // 这里会等待上面的goroutine完成
	// }()

	// time.Sleep(500 * time.Millisecond) // 确保上面的Wait开始执行
	// wg.Add(1)                          // 这里尝试在Wait执行的同时增加计数，可能触发panic
	// go func() {
	// 	time.Sleep(time.Second) // 模拟一些工作
	// 	wg.Done()
	// }()
	// wg.Wait() // 等待所有任务完成
}
