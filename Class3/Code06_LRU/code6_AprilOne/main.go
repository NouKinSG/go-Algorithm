package main

import "sync"

type LRUCache struct {
	// 存东西
	cache    map[int]*DLinkList
	size     int
	capacity int
	head     *DLinkList
	tail     *DLinkList
	mu       sync.Mutex
}

func Constructor(capacity int) LRUCache {
	head := &DLinkList{}
	tail := &DLinkList{}

	head.Next = tail
	tail.Pre = head
	return LRUCache{
		cache:    make(map[int]*DLinkList),
		capacity: capacity,
		head:     head,
		tail:     tail,
		mu:       sync.Mutex{},
	}
}

func (this *LRUCache) Get(key int) int {
	this.mu.Lock()         // 加锁
	defer this.mu.Unlock() // 释放锁
	if node, ok := this.cache[key]; ok {
		// 使用了，更新
		this.moveToHead(node)
		return node.Val
	}

	// 没有
	return -1
}

func (this *LRUCache) Put(key int, value int) {

	this.mu.Lock()

	defer this.mu.Unlock()

	// 加入前看看有没有
	if node, ok := this.cache[key]; ok {
		// 有，更新val
		node.Val = value
		// 更新使用
		this.moveToHead(node)
		return
	}

	// 没有，加入操作
	NewNode := &DLinkList{Key: key, Val: value}

	if this.size == this.capacity {
		// 缓冲区满了
		// 先删除，再加入
		this.removeLRUItem()
		this.size--
	}

	this.cache[key] = NewNode
	// 更新使用
	this.AddToHead(NewNode)
	this.size++
}

type DLinkList struct {
	Key, Val int
	Pre      *DLinkList
	Next     *DLinkList
}

// 删除某一节点
func (this *LRUCache) removeNode(node *DLinkList) {
	node.Next.Pre = node.Pre
	node.Pre.Next = node.Next
}

// 头部添加
func (this *LRUCache) AddToHead(node *DLinkList) {
	node.Next = this.head.Next
	node.Pre = this.head

	this.head.Next.Pre = node
	this.head.Next = node
}

// 将某个节点移动至头部
func (this *LRUCache) moveToHead(node *DLinkList) {
	this.removeNode(node)
	this.AddToHead(node)
}

// 队尾
func (this *LRUCache) removeLRUItem() {
	tail := this.tail.Pre
	this.removeNode(tail)
	delete(this.cache, tail.Key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
