package GMP

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"testing"
	"time"
)

var (
	// 状态计数器
	goroutineCount = 0
	mu             sync.Mutex
)

// 不是 goroutine 进行访问时不需要加锁
func Add() {
	goroutineCount++
}

// goroutine 并发访问的变量，需要加锁
func Exit() {
	mu.Lock()
	defer mu.Unlock()
	goroutineCount--
}

var (
	wg sync.WaitGroup
)

func TestRunTask(t *testing.T) {
	syncRun()
	wg.Wait()
}

func runTask(id int) {
	// 结束时，计数器 - 1
	defer wg.Done()

	fmt.Printf("task %d start..\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("task %d complete\n", id)
}

func syncRun() {
	for i := 0; i < 10; i++ {
		go runTask(i + 1)
		// 每启动一个 goroutine 就+1
		wg.Add(1)
	}
}

func asyncRun() {
	for i := 0; i < 10; i++ {
		go runTask(i + 1)
		// 没启动一个go routine 就+1
		wg.Add(1)
	}
}

func TestTrace(t *testing.T) {
	// 创建 trace 文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// 启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	defer trace.Stop()

	asyncRun()
	wg.Wait()
}
