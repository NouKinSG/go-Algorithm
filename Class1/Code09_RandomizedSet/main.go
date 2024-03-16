package main

import (
	"math/rand"
	"time"
)

type RandommizedSet struct {
	// 插入、删除、随机访问.插入、删除可以用哈希表，而随机访问只能用 数组

	//哈希表，用于插入、删除
	indexMap map[int]int

	//随机访问用 数组（切片）
	values []int

	// 随机数生成器 字段
	rng *rand.Rand
}

func Constructor() RandommizedSet {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	return RandommizedSet{
		indexMap: make(map[int]int),
		values:   []int{},
		// 等价于  <=> value:   make([]int, 0),
		rng: rng,
	}
}

func (this *RandommizedSet) Insert(val int) bool {
	// 已经有了，返回false
	if _, ok := this.indexMap[val]; ok {
		return false
	}

	// 无元素，做添加操作
	this.indexMap[val] = len(this.values)
	this.values = append(this.values, val)
	return true
}

func (this *RandommizedSet) Remove(val int) bool {
	//  没有 返回false
	index, ok := this.indexMap[val]
	if !ok {
		return false // val不存在，返回false
	}

	// 下面是 有的情况：

	// 数组中最后一个元素的索引
	lastIndex := len(this.values) - 1
	// 数组中最后一个元素对应的值
	lastValue := this.values[lastIndex]

	// 补空,最后一个元素的位置，放入删除位置
	this.values[index] = lastValue
	this.indexMap[lastValue] = index //重新绑定映射

	delete(this.indexMap, val)            // 删除 val映射
	this.values = this.values[:lastIndex] // 删除 数组中最后一个元素

	return true
}

func (this *RandommizedSet) GetRandom() int {
	randomIndex := this.rng.Intn(len(this.values))
	return this.values[randomIndex]
}

func main() {
	// Your code here
}
