package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

// 增，尾部增加
func (this *LinkedList) append(value int) {
	// 构造节点
	newNode := &ListNode{val: value}

	if this.head == nil {
		this.head = newNode
		return
	}
	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = newNode
}

// 删   删除链表中的一个元素
func (this *LinkedList) Remove(value int) {
	if this.head == nil {
		return
	}
	// 否则，找有没有 值为value的节点
	if this.head.val == value {
		this.head = this.head.next
		return
	}
	cur := this.head
	for cur.next != nil {
		if cur.next.val == value {
			cur.next = cur.next.next
			return
		}
		cur = cur.next
	}
}

// 改
func (this *LinkedList) Update(oldVal, newVal int) {
	cur := this.head
	for cur != nil {
		if cur.val == oldVal {
			cur.val = newVal
			return
		}
		cur = cur.next
	}
}

// 查
func (this *LinkedList) Find(value int) bool {
	cur := this.head
	for cur != nil {
		if cur.val == value {
			return true
		}
		cur = cur.next
	}
	return false
}

func (this *LinkedList) Print() {
	cur := this.head
	for cur != nil {

		fmt.Printf("%d -> ", cur.val)
		cur = cur.next
	}
	fmt.Println(nil)
}
