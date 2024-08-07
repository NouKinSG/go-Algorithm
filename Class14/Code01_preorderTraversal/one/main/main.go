package main

import (
	"fmt"
	"sync"
)

func generator(nums []int, out chan<- int) {
	for _, n := range nums {
		out <- n
	}
	close(out)
}

func processor(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range in {
		out <- n * 2
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	genChan := make(chan int)
	procChan := make(chan int)
	var wg sync.WaitGroup

	go generator(nums, genChan)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go processor(genChan, procChan, &wg)
	}

	go func() {
		wg.Wait()
		close(procChan)
	}()

	for result := range procChan {
		fmt.Println(result)
	}
}
