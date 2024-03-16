package main

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

func main() {
	//创建链表

}
