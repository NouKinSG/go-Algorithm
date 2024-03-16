package main

import "unsafe"

const DEFAULT = false
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func Constant() {
	// Your code goes here
	println(a, b, c)
}
