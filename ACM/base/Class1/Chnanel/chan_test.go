package Chnanel

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
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
	ch := make(chan string, 5) // 创建一个无缓冲通道

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

func A(startA, startB chan struct{}) {
	fmt.Println("AAAAAAAAAAAAAAAAAAAA")
	a := []string{"1", "2", "3"}
	index := 0
	for range startA {
		if index > 2 {
			return
		}
		fmt.Println(a[index])
		index++
		startB <- struct{}{}
	}
}

func B(startA, startB chan struct{}) {
	fmt.Println("BBBBBBBBBBBBBBB")
	b := []string{"x", "y", "z"}
	index := 0
	for range startB {
		fmt.Println(b[index])
		index++
		startA <- struct{}{}
	}
}

func TestAltanation(t *testing.T) {
	startA, startB := make(chan struct{}), make(chan struct{})

	go A(startA, startB)
	go B(startA, startB)

	//startA <- struct{}{}
	fmt.Println("mainmainmainmain")

}

func TestDeadLock(t *testing.T) {
	ch := make(chan string)
	// send

	// send in a new goroutine
	go func() {
		ch <- "hello"
	}()

	// receive
	fmt.Println(<-ch)

}

func TestDeadLockV2(t *testing.T) {
	ch := make(chan string, 1)
	// send
	{
		ch <- "hello"
	}

	// receive
	{
		fmt.Println(<-ch)
	}
}

func senderV2(ch chan string, down chan struct{}) {
	ch <- "hello"
	ch <- "this"
	ch <- "is"
	ch <- "alice"
	// 发送通话结束
	ch <- "EOF"

	// 同步模式等待recver 处理完成
	<-down
	// 处理完成后关闭channel
	close(ch)
}

// recver 循环读取chan里面的数据，直到channel关闭
func recverV2(ch chan string, down chan struct{}) {
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

func TestBufferedChan(t *testing.T) {
	ch := make(chan string, 5)

	down := make(chan struct{})
	go senderV2(ch, down) // sender goroutine
	go recverV2(ch, down) // recver goroutine

	<-down
}

type Task struct {
	ID         int
	JobID      int
	Status     string
	CreateTime time.Time
}

func (t *Task) Run() {
	sleep := rand.Intn(1000)
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	t.Status = "Completed"
}

var wg sync.WaitGroup

// worker的数量，即使用多少goroutine执行任务
const workerNum = 3

func RunTaskWithPool() {
	wg.Add(workerNum)

	// 创建容量为10的buffered channel
	taskQueue := make(chan *Task, 10)

	// 激活goroutine，执行任务
	for workID := 0; workID <= workerNum; workID++ {
		go worker(taskQueue, workID)
	}

	// 将待执行任务放进buffered channel，共15个任务
	for i := 1; i <= 15; i++ {
		taskQueue <- &Task{
			ID:         i,
			JobID:      100 + i,
			CreateTime: time.Now(),
		}
	}

	wg.Wait()

	//记得关闭channel
	close(taskQueue)
}

// 从buffered channel中读取任务，并执行任务
func worker(in chan *Task, workID int) {
	defer wg.Done()
	for v := range in {
		fmt.Printf("Worker%d: recv a request: TaskID:%d, JobID:%d\n", workID, v.ID, v.JobID)
		v.Run()
		fmt.Printf("Worker%d: Completed for TaskID:%d, JobID:%d\n", workID, v.ID, v.JobID)
	}
}
