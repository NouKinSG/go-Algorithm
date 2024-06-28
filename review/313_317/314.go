package _13_317

import (
	"math/rand"
	"time"
)

// 3.14 O(1)
// 380. O(1) 时间插入、删除和获取随机元素

type RandomizedSet struct {
	// O(1)存取 用map
	indexMap map[int]int
	// 随机存取用数组
	arrSlice []int
	// 随机数种子
	rng *rand.Rand
}

func Constructor() RandomizedSet {
	source := rand.NewSource(time.Now().Unix())
	rng := rand.New(source)

	return RandomizedSet{
		indexMap: make(map[int]int),
		arrSlice: []int{},
		rng:      rng,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.indexMap[val]
	if ok {
		return false
	}
	// 不存在,执行插入
	this.indexMap[val] = len(this.arrSlice)
	this.arrSlice = append(this.arrSlice, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	// 删除前看看元素存不存在
	oldIndex, ok := this.indexMap[val]
	if !ok {
		// 不存在直接返回false
		return false
	}
	// 存在
	// 执行删除
	lastIndex := len(this.arrSlice) - 1
	lastValue := this.arrSlice[lastIndex]

	// 添加新映射
	this.indexMap[lastValue] = oldIndex
	this.arrSlice[oldIndex] = lastValue

	// 删除旧映射
	this.arrSlice = this.arrSlice[:lastIndex]
	delete(this.indexMap, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	randIndex := this.rng.Intn(len(this.arrSlice))
	return this.arrSlice[randIndex]
}
