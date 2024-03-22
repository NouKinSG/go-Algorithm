package main

import (
	"math/rand"
	"time"
)

type RandomizedSet4 struct {
	// 删改查 用 map
	indexMap map[int]int

	// 随机存取用 arr
	arrSlice []int

	// 随机数生成器
	rng *rand.Rand
}

func Constructor4() RandomizedSet4 {
	src := rand.NewSource(time.Now().UnixNano())
	rg := rand.New(src)
	return RandomizedSet4{
		indexMap: make(map[int]int),
		arrSlice: []int{},
		rng:      rg,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	// 插入前看看存不存在
	if _, ok := this.indexMap[val]; ok {
		return false
	}

	// 不存在，执行插入
	this.indexMap[val] = len(this.arrSlices)
	this.arrSlices = append(this.arrSlices, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	// 删除前，看看存不存在
	index, ok := this.indexMap[val]
	if !ok {
		return false
	}

	// 存在，需要删除操作
	// 删除步骤 ——> 数组的最后一个元素，替换删除元素，添加新映射，删除旧映射，删除最后一个元素
	lastIndex := len(this.arrSlices) - 1
	lastValue := this.arrSlices[lastIndex]

	// 替换
	this.arrSlices[index] = lastValue
	this.indexMap[lastValue] = index

	// 删除
	delete(this.indexMap, val)
	this.arrSlices = this.arrSlices[:lastIndex]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	randIndex := this.rng.Intn(len(this.arrSlices))
	return this.arrSlices[randIndex]
}
