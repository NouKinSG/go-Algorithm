package test

import (
	"fmt"
	"testing"
)

type Node struct {
	val  int
	next *Node
}

func reverseList(head *Node) *Node {
	var pre *Node
	var next *Node

	for head != nil {
		next = head.next
		head.next = pre
		pre = head
		head = next
	}
	return pre
}

func TestReverseList(t *testing.T) {
	fmt.Println("test")
	var emptyList *Node

	// 测试用例 空链表
	reversedEmptyList := reverseList(emptyList)
	if reversedEmptyList != nil {
		t.Errorf("TestReverseList failed, got %v", reversedEmptyList)
	}

}
