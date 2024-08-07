package one

// 前缀树节点
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// 整颗前缀树
type Trie struct {
	root *TrieNode
}

// Constructor 初始化前缀树对象
func Constructor() Trie {
	// 初始化 根节点
	return Trie{
		root: &TrieNode{children: make(map[rune]*TrieNode)},
	}
}

// Insert 向前缀树中插入字符串
func (this *Trie) Insert(word string) {
	node := this.root // 从根节点开始
	for _, char := range word {
		// 如果不在前缀树中才加入
		if _, ok := node.children[char]; !ok {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		// 移动到子节点
		node = node.children[char]
	}
	node.isEnd = true // 最后一个字符节点标记为结束
}

// Search 检查某个字符串是否在前缀树中
func (this *Trie) Search(word string) bool {
	// 从跟节点开始
	node := this.root
	for _, char := range word {
		// 不在前缀树中
		if _, ok := node.children[char]; !ok {
			return false
		}
		// 移动到子节点
		node = node.children[char]
	}
	// 说明整个word都在树中，但是要看看是不是单词结束
	return node.isEnd
}

// StartsWith 检查是否存在前缀 prefix 开头的单词
func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, char := range prefix {
		// 如果不存在，直接返回false
		if _, ok := node.children[char]; !ok {
			// 直接返回false
			return false
		}
		// 移动到子节点
		node = node.children[char]
	}
	// 如果没有返回false，这里直接返回true，因为不需要看有没有到末尾
	return true
}
