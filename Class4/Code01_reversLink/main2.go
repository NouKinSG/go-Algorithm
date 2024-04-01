package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}
	return pre
}

type DoubleListNode struct {
	Val  int
	Next *DoubleListNode
	Pre  *DoubleListNode
}

func reverseDoubleList(head *DoubleListNode) *DoubleListNode {
	// pre 和 cur
	var pre *DoubleListNode = nil
	cur := head

	for cur != nil {
		// 先拿 next节点
		next := cur.Next
		// 处理当前节点
		cur.Next = pre
		cur.Pre = next

		// 处理完当前处理pre节点
		// 判空
		if pre != nil {
			pre.Pre = cur
		}

		pre = cur
		cur = next
	}
	return pre
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}
func quick(arr []int, left, right int) {
	if left < right {
		privot := part(arr, left, right)
		quick(arr, left, privot-1)
		quick(arr, privot+1, right)
	}
}

func part(arr []int, left, right int) int {
	privot := arr[right]
	index := left - 1
	for j := left; j < right; j++ {
		if arr[j] < privot {
			index++
			arr[index], arr[j] = arr[j], arr[index]
		}
	}
	arr[index+1], arr[right] = arr[right], arr[index+1]
	return index + 1
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{9, 5, 1, 8, 3, 7, 4, 6, 2}
	fmt.Println("Before sorting:", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("After sorting:", arr)
}
