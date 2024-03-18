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

type RandomizedSet struct {
	// 增删改查用 map
	indexMap map[int]int
	// 随机存取用 数组(切片)
	arrSlices []int
	//随机数生成器
	rng *rand.Rand
}

// 构造函数
func NewRandomizedSet() *RandomizedSet {
	src := rand.NewSource(time.Now().UnixNano())
	rg := rand.New(src)
	return &RandomizedSet{
		indexMap:  make(map[int]int),
		arrSlices: []int{},
		rng:       rg,
	}
}

// 插入
func (this *RandomizedSet) Insert(val int) bool {
	// 判断有没有，没有再插入
	if _, ok := this.indexMap[val]; ok {
		return false
	}

	// 没有 执行插入
	this.indexMap[val] = len(this.arrSlices)
	this.arrSlices = append(this.arrSlices, val)
	return true
}

// 删除
func (this *RandomizedSet) Remove(val int) bool {
	// 先看看有没有元素，有才删除
	index, ok := this.indexMap[val]
	if !ok {
		return false
	}
	// 有，执行删除
	// 删除步骤，用数组最后一个元素，替换当前元素位置，添加新映射。删除最后一个元素，删除旧映射

	// 获取最后一个元素信息
	lastIndex := len(this.arrSlices) - 1
	lastValue := this.arrSlices[lastIndex]

	this.arrSlices[index] = lastValue
	this.indexMap[lastValue] = index
	delete(this.indexMap, val)
	this.arrSlices = this.arrSlices[:lastIndex]
	return true
}

// 随机取数
func (this *RandomizedSet) GetRandom() int {
	randIndex := this.rng.Intn(len(this.arrSlices))
	return this.arrSlices[randIndex]
}
