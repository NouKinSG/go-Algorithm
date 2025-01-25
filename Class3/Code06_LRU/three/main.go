package three

type LRUCache struct {
	// 容量大小
	capacity int

	// 当前大小
	size int
	// 哈希表，存储 key 和 节点 映射关系
	cache map[int]*RListNode

	// 双向链表
	head *RListNode
	tail *RListNode
}

func Constructor(capacity int) LRUCache {
	// 初始化双向链表
	head := &RListNode{}
	tail := &RListNode{}

	head.next = tail
	tail.pre = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*RListNode),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	// Get的时候看看有没有
	if node, ok := this.cache[key]; ok {
		// 有，移动至最新
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 放置前，看看有没有
	if node, ok := this.cache[key]; ok {
		// 有，将 key 和 value 重新绑定,更新最近使用即可
		node.val = value
		this.moveToHead(node)
		return
	}

	// 没有，先看看有没有满
	if this.size == this.capacity {
		// 满了，删除一个
		this.removeLRUItem()
		this.size--
	}
	NewNode := &RListNode{key: key, val: value}
	// 绑定映射
	this.cache[key] = NewNode
	// 添加至最新使用
	this.addToHead(NewNode)
	this.size++
}

// 使用双向链表
type RListNode struct {
	key  int
	val  int
	pre  *RListNode
	next *RListNode
}

// 双向链表，删除节点
func (this *LRUCache) removeNode(node *RListNode) {
	//  A ⇄ B ⇄ C，假如删除 B
	// 需要把 A 的下一个 指向 C
	node.pre.next = node.next

	// 再把 C 的上一个指向 A
	node.next.pre = node.pre
}

// 从头节点插入链表
func (this *LRUCache) addToHead(node *RListNode) {
	// head ⇄ A ⇄ B ⇄ C ⇄ tail，假如想插入一个 O 在头节点，变成 head ⇄ O ⇄ A ⇄ B ⇄ C ⇄ tail
	node.pre = this.head
	node.next = this.head.next

	//	处理head
	this.head.next.pre = node
	this.head.next = node
}

// 将节点移动至 头部
func (this *LRUCache) moveToHead(node *RListNode) {
	this.removeNode(node)
	this.addToHead(node)
}

// 删除最近最久未使用
func (this *LRUCache) removeLRUItem() {
	lru := this.tail.pre
	this.removeNode(lru)
	delete(this.cache, lru.key)
}
