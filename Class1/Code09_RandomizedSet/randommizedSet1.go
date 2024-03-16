package main

import (
	"math/rand"
	"time"
)

type RandommizedSet1 struct {
	// 增删改查 用map
	indexMap map[int]int

	// 随机存取 用数组(切片)
	arrSlice []int

	// 本地随机数生成器
	rng *rand.Rand
}

// 构造函数
func NewRandommisedSet1() RandommizedSet1 {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	return RandommizedSet1{
		indexMap: make(map[int]int),
		arrSlice: []int{},
		rng:      rnd,
	}
}

func (this *RandommizedSet1) Insert(val int) bool {
	// 如果已经有了，返回false
	if _, ok := this.indexMap[val]; ok {
		return false
	}

	// 没有，进行插入操作
	this.indexMap[val] = len(this.arrSlice) // 存索引必须先 建立映射再，数组添加
	this.arrSlice = append(this.arrSlice, val)

	return true
}

func (this *RandommizedSet1) Remove(val int) bool {
	index, ok := this.indexMap[val]
	// 如果没有，则无法删除。返回false
	if !ok {
		return false
	}

	// 一下是有，做删除操作
	// 如何删除，将数组中删除的元素 被 最后一个元素补进来。删除最后一个元素,重新绑定映射，删除最后映射
	lastIndex := len(this.arrSlice) - 1
	lastValue := this.arrSlice[lastIndex]

	// 替换元素
	this.arrSlice[index] = lastValue
	// 绑定映射
	this.indexMap[lastValue] = index
	// 删除最后元素
	this.arrSlice = this.arrSlice[:lastIndex]
	// 删除旧映射
	delete(this.indexMap, val)

	return true
}

func (this *RandommizedSet1) GetRandom() int {
	index := this.rng.Intn(len(this.arrSlice))
	return this.arrSlice[index]
}
