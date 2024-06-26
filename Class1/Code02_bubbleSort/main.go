package main

func main() {
	// go say("world")
	// say("hello")
	// gorountinemain()

	text := "Hello, world!"
	done := make(chan bool)

	go typewriter(text, done)

	// 等待打字机完成
	<-done
}

func Bubble(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}
func Bubble2(arr []int) {
	if len(arr) < 2 {
		return
	}
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
