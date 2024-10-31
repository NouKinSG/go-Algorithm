package main

import "fmt"

// ListNode 定义链表结构
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 合并链表的公共前缀和公共后缀部分
 * @param a ListNode类 链表a
 * @param b ListNode类 链表b
 * @return ListNode类 合并后的链表
 */
func mergeList(a *ListNode, b *ListNode) *ListNode {
	// 提取公共前缀，并同时更新 a 和 b 的指针位置
	prefix, aNew, bNew := extractPrefix(a, b)
	// 获取链表长度
	aLen := getLength(aNew)
	bLen := getLength(bNew)

	// 对齐两个链表的起始点
	if aLen > bLen {
		aNew = advanceBy(aNew, aLen-bLen)
	} else if bLen > aLen {
		bNew = advanceBy(bNew, bLen-aLen)
	}

	// 找到公共后缀起点（基于节点值比较）
	common := findCommonSuffix(aNew, bNew)

	// 拼接前缀和后缀
	if prefix != nil {
		last := prefix
		for last.Next != nil {
			last = last.Next
		}
		last.Next = common // 拼接到公共后缀
		return prefix
	}
	return common
}

// 获取链表长度
func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

// 将链表前进指定步数
func advanceBy(head *ListNode, steps int) *ListNode {
	for steps > 0 && head != nil {
		head = head.Next
		steps--
	}
	return head
}

// 提取链表的公共前缀部分，同时返回更新后的 a 和 b
func extractPrefix(a, b *ListNode) (*ListNode, *ListNode, *ListNode) {
	dummy := &ListNode{}
	cur := dummy
	for a != nil && b != nil && a.Val == b.Val {
		cur.Next = &ListNode{Val: a.Val}
		cur = cur.Next
		a = a.Next
		b = b.Next
	}
	return dummy.Next, a, b
}

// 查找两个链表的公共后缀部分，基于节点值比较
func findCommonSuffix(a, b *ListNode) *ListNode {
	// 使用两个栈存储链表a和b的节点值
	stackA, stackB := []int{}, []int{}
	for a != nil {
		stackA = append(stackA, a.Val)
		a = a.Next
	}
	for b != nil {
		stackB = append(stackB, b.Val)
		b = b.Next
	}

	// 通过比较栈顶元素，查找公共后缀
	var common *ListNode
	for len(stackA) > 0 && len(stackB) > 0 {
		valA := stackA[len(stackA)-1]
		valB := stackB[len(stackB)-1]
		if valA == valB {
			// 如果值相同，继续寻找公共后缀
			common = &ListNode{Val: valA, Next: common}
			stackA = stackA[:len(stackA)-1]
			stackB = stackB[:len(stackB)-1]
		} else {
			break
		}
	}
	return common
}

// 打印链表
func printList(head *ListNode) {
	if head == nil {
		fmt.Println("nil")
		return
	}
	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Printf(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// 测试用例: 输入两个链表
	a := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	b := &ListNode{1, &ListNode{2, &ListNode{4, &ListNode{5, nil}}}}

	// 合并链表
	result := mergeList(a, b)

	// 输出合并后的链表
	fmt.Print("合并后的链表: ")
	printList(result) // 预期输出: 1 -> 2 -> 3 -> 4
}
