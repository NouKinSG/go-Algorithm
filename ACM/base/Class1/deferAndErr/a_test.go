package deferAndErr

import (
	"fmt"
	"testing"
)

func TestDEFER(t *testing.T) {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
	fmt.Println("end")
}

func TestA(t *testing.T) {
	fmt.Println("func start")
	x := 10
	defer func(x int) {
		fmt.Println("in defer: ", x)
	}(x)
	x = 30
	fmt.Println("func end: ", x)
}

func TestB(t *testing.T) {
	fmt.Println("func start")
	x := 10
	defer func(x *int) {
		fmt.Println("in defer: ", *x)
	}(&x)
	x = 30
	fmt.Println("func end: ", x)
}

// defer与闭包
func TestDeferAndClosure(t *testing.T) {
	fmt.Println("func start")
	x := 10
	defer func() {
		fmt.Println("in defer:", x)
	}()
	x = 30
	fmt.Println("func end:", x)
}
