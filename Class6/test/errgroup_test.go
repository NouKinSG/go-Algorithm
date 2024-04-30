package test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

func TestErr(t *testing.T) {
	var g errgroup.Group

	for i := 0; i < 10; i++ {
		// i := i // 注意：捕获迭代变量
		g.Go(func() error {
			// 执行任务，如 HTTP 请求、数据库操作等
			if i == 5 { // 假设在 i == 5 时任务失败
				return fmt.Errorf("error at i=%d", i)
			}
			fmt.Println("Processing:", i)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("Received error:", err)
	} else {
		fmt.Println("All tasks completed successfully")
	}
}

// 模拟无限重试，协程 “阻塞” 问题
func TestGoroutine(t *testing.T) {

	var ch chan int
	quit := make(chan struct{})

	start := func(idx int) {
		t.Logf("start: %d", idx)
		ch = make(chan int, 100000)
		for {
			inner := 0
			select {
			case <-quit:
				t.Logf("退出了 start ch：%d", idx)
				return
			case num := <-ch:
				inner++
				// 模拟去进行前置获取
				for {
					nowStr := time.Now().Format("15:04:05")
					// 假设失败了，一直重试
					t.Logf("[%s] 获取第 %d 批，消费第 %d 个内容: %d", nowStr, idx, inner, num)
					time.Sleep(time.Second)
				}
			case <-time.After(5 * time.Second):
				// 超时控制
				t.Logf("超时，退出：%d", idx)
			}
		}

	}

	// 模拟重新去获取对应的操作。
	restart := func(idx int) {
		t.Logf("restart: %d", idx)
		close(quit)
		quit = make(chan struct{})
		start(idx)
	}

	// 等待这个生产协程退出，就退出整个程序
	var wg sync.WaitGroup
	wg.Add(1)

	// 模拟一分钟，自行观察日志
	producer := func() {
		defer wg.Done()
		// 生产内容，每隔一秒就生成一个消息
		for i := 1; i <= 60; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
	}

	// 消费
	go start(1)
	// 确保 ch 准备好了
	time.Sleep(time.Millisecond * 500)

	// 生产
	go producer()

	// 模拟修改了三次配置
	time.Sleep(3 * time.Second)
	go restart(2)
	time.Sleep(3 * time.Second)
	go restart(3)
	time.Sleep(3 * time.Second)
	go restart(4)

	// 等待生产数据的协程退出
	wg.Wait()
}
