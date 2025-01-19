package two

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	// 存取用 map
	indexMap map[int]int
	// 随机存取用 数组
	SliceArray []int
	// 随机数种子
	rand *rand.Rand
}

func Constructor() RandomizedSet {
	// 生成随机数种子
	src := rand.NewSource(time.Now().UnixNano())
	rd := rand.New(src)

	// 返回
	return RandomizedSet{
		indexMap:   make(map[int]int),
		SliceArray: []int{},
		rand:       rd,
	}
}

// Insert  当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
func (this *RandomizedSet) Insert(val int) bool {
	// 插入前，看看元素有没有
	if _, ok := this.indexMap[val]; ok {
		// 存在了
		return false
	}

	// 不存在，执行插入
	this.indexMap[val] = len(this.SliceArray)
	this.SliceArray = append(this.SliceArray, val)
	return true
}

// Remove bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
func (this *RandomizedSet) Remove(val int) bool {
	// 删除前看看有没有
	oldIndex, ok := this.indexMap[val]
	if !ok {
		// 没有，直接返回
		return false
	}

	// 存在
	// 执行删除
	// 如何删除，找到映射
	// 1.最后一个元素的 index和value
	lastIndex := len(this.SliceArray) - 1
	lastValue := this.SliceArray[lastIndex]

	// 2.需要删除元素的index和下标，被替换掉
	this.SliceArray[oldIndex] = lastValue
	this.indexMap[lastValue] = oldIndex

	this.SliceArray = this.SliceArray[:lastIndex]
	// 删除映射
	delete(this.indexMap, val)
	return true
}

// GetRandom 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
func (this *RandomizedSet) GetRandom() int {
	randIndex := this.rand.Intn(len(this.SliceArray))
	return this.SliceArray[randIndex]
}
