package _map

import (
	"fmt"
	"testing"
)

func TestHMap(t *testing.T) {
	var m map[string]int // *hmap
	fmt.Printf("%p\n", m)
}

func TestHMapMake(t *testing.T) {
	var m map[string]int
	m = make(map[string]int)
	fmt.Printf("%p\n", m)
}

func TestHMapInit(t *testing.T) {
	var a map[int]string
	var b []string
	c := make(map[int]int)
	fmt.Printf("%p,%p, %p\n", a, b, c)
	//a[0] = "a"
	//b[0] = "a"

	//a = map[int]string{0: "a"}
	//b = []string{"a"}
	//fmt.Printf("%p, %p\n", a, b)
	a = map[int]string{}
	b = []string{}
	fmt.Printf("%p, %p\n", a, b)
}
