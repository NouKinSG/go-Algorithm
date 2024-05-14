package Btest

type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表基础题目
// 1. 反转链表 https://leetcode-cn.com/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 2. 单链表的中间节点 https://leetcode-cn.com/problems/middle-of-the-linked-list/
func middleList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 3. 合并两个有序链表 https://leetcode-cn.com/problems/merge-two-sorted-lists/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 前置校验 || 递归终止条件
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 过程处理
	var res *ListNode
	if l1.Val < l2.Val {
		res = l1
		res.Next = mergeTwoLists(l1.Next, l2)
	}
	if l1.Val >= l2.Val {
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)
	}
	return res
}

// 4. 删除链表的倒数第 N 个结点 https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	count := 0
	for cur != nil {
		cur = cur.Next
		count++
	}
	m := count - n
	cur = head
	if m == 0 {
		return head.Next
	}
	for i := 1; i <= m; i++ {
		if i == m {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return head
}

// 5. 链表排序 https://leetcode-cn.com/problems/sort-list/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 不允许使用前面的函数1，2，3，4那些函数，需要自己写
	// 1. 找到链表的中间节点
	// 2. 递归拆分链表
	// 3. 合并链表
	// 4. 返回合并后的链表
	mid := middleList(head)
	left := head
	right := mid.Next
	mid.Next = nil
	return mergeTwoLists(sortList(left), sortList(right))
}

// 双向链表
type DoubleListNode struct {
	Val  int
	Key  int
	Pre  *DoubleListNode
	Next *DoubleListNode
}

// 6. LRU缓存机制 https://leetcode-cn.com/problems/lru-cache/
type LRUCache struct {
	Capacity int
	Length   int
	Head     *DoubleListNode
	Tail     *DoubleListNode
	Hash     map[int]*DoubleListNode
}

func Constructor(capacity int) LRUCache {
	head, tail := &DoubleListNode{}, &DoubleListNode{}
	head.Next = tail
	tail.Pre = head
	return LRUCache{
		capacity,
		0,
		head,
		tail,
		make(map[int]*DoubleListNode),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Hash[key]; ok {
		// 更新使用
		this.moveToHead(node)
		return node.Val
	}
	// 没有
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Hash[key]; ok {
		//更新
		node.Val = value
		this.moveToHead(node)
		return
	}
	// 检查size
	if this.Length >= this.Capacity {
		this.removeLRUItem()
		this.Length--
	}

	// 没有
	node := &DoubleListNode{
		Val: value,
		Key: key,
	}
	this.addToHead(node)
	this.Hash[key] = node
	this.Length++
}

// 双向链表操作
// 删除节点
func (this *LRUCache) removeNode(node *DoubleListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

// 在头部添加节点
func (this *LRUCache) addToHead(node *DoubleListNode) {
	node.Next = this.Head.Next
	node.Pre = this.Head

	this.Head.Next.Pre = node
	this.Head.Next = node
}

// 移动节点
func (this *LRUCache) moveToHead(node *DoubleListNode) {
	this.removeNode(node)
	this.addToHead(node)
}

// 删除
func (this *LRUCache) removeLRUItem() {
	tail := this.Tail.Pre
	this.removeNode(tail)
	delete(this.Hash, tail.Key)
}
