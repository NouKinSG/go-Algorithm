package main

type RLinkNode1 struct {
	key, val int
	pre      *RLinkNode1
	next     *RLinkNode1
}

type LRUCache1 struct {
	// 存储数据
	cache map[int]*RLinkNode1

	// 容量
	capacity int

	// 当前容量
	size int

	// 链表格式
	head *RLinkNode1
	tail *RLinkNode1
}

func Constructor1(capacity int) LRUCache1 {
	head := &RLinkNode1{}
	tail := &RLinkNode1{}

	head.next = tail
	tail.pre = head
	return LRUCache1{
		cache:    make(map[int]*RLinkNode1),
		capacity: capacity,
		size:     0,
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache1) Get(key int) int {
	// 看看有没有
	if node, ok := this.cache[key]; ok {
		// 移动至 最近使用 ——> 头节点
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache1) Put(key int, value int) {
	// put前 看看有没有
	if node, ok := this.cache[key]; ok {
		// 有了,更新值
		node.val = value
		this.moveToHead(node)
		return
	}

	// 没有 做put操作，但是先看看容量有没有满
	if this.size == this.capacity {
		this.removeLRUItem()
		this.size--
	}

	NewNode := &RLinkNode1{key: key, val: value}
	this.cache[key] = NewNode
	this.moveToHead(NewNode)
	this.size--
}

func (this *LRUCache1) removeNode(node *RLinkNode1) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

// 头插入 节点
func (this *LRUCache1) addToHead(node *RLinkNode1) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node
}

// 将一个节点移动至头部
func (this *LRUCache1) moveToHead(node *RLinkNode1) {
	this.removeNode(node)
	this.addToHead(node)
}

// 最近最久未使用
func (this *LRUCache1) removeLRUItem() {
	tail := this.tail.pre
	this.removeNode(tail)
	delete(this.cache, tail.key)
}
