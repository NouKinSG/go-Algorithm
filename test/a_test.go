package test

import (
	"fmt"
	"testing"
)

// 二分查找
func findtwo(nums []int, target int) int {
	ans := 0
	if len(nums) == 0 {
		return ans
	}
	ans = find(nums, target)
	return ans
}

// 返回索引
func find(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left++
		} else {
			right--
		}
	}
	return -1
}

func TestFind(t *testing.T) {
	//arr := []int{1, 3, 5, 7, 9, 10}
	//index := findtwo(arr, 10)
	//fmt.Println(index)

	head := &ListNode{val: 3, next: &ListNode{val: 2, next: &ListNode{val: 1, next: nil}}}

	ans := reverse(head)
	fmt.Println(ans.val)
	fmt.Println(ans.next.val)
	fmt.Println(ans.next.next.val)
}

type ListNode struct {
	val  int
	next *ListNode
}

// 反转链表
func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.next
		cur.next = pre
		pre = cur
		cur = next
	}
	return pre
}
