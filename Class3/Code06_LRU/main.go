package main

// 双向链表
type RListNode struct {
	key  int
	val  int
	pre  *RListNode
	next *RListNode
}

type LRUCache struct {
	// 容量大小
	capacity int

	// 当前容量
	size int

	// 哈希表，存储 key 和 节点的映射关系
	cache map[int]*RListNode

	// 链表格式，头为节点
	head *RListNode
	tail *RListNode
}

// 初始化 LRU
func Constructor(capacity int) *LRUCache {
	head := &RListNode{}
	tail := &RListNode{}

	head.next = tail
	tail.pre = head

	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*RListNode),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	// 看看有没有
	if node, ok := this.cache[key]; ok {
		// 移至最新使用 head
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// put 之前看看有没有
	if node, ok := this.cache[key]; ok {
		// 如果已经有了
		node.val = value
		this.moveToHead(node)
		return
	}
	// 没有需要put，put前再看看有没有缓存有没有满
	if this.size == this.capacity {
		// 满了，需要删除 最近最久未使用节点
		this.removeLRUItem()
		this.size--
	}

	newNode := &RListNode{key: key, val: value}
	// 绑定映射
	this.cache[key] = newNode
	// 添加至 最新使用 head
	this.addToHead(newNode)
	this.size++
}

// 双向链表的删除节点
func (this *LRUCache) removeNode(node *RListNode) {
	// 当前节点的，前一个节点，指向当前节点的下一个
	node.pre.next = node.next
	// 同时 下一个节点的pre，指向当前节点的pre
	node.next.pre = node.pre
}

// 从头插入节点
func (this *LRUCache) addToHead(node *RListNode) {
	// 从头插入链表

	// 处理 node
	// node 的前节点等于 head
	node.pre = this.head
	// 自己的下一个等于 链表头的下一个
	node.next = this.head.next

	// 处理 head
	this.head.next.pre = node
	this.head.next = node
}

// 将节点移动至头
func (this *LRUCache) moveToHead(node *RListNode) {
	// 先删除 该节点
	this.removeNode(node)
	// 在头部添加
	this.addToHead(node)
}

// 删除 最近最久未使用
func (this *LRUCache) removeLRUItem() {
	lru := this.tail.pre
	this.removeNode(lru)
	delete(this.cache, lru.key)
}
