package main

sync.WaitGroup{}

func numIslands(grid [][]byte) int {
	ans := 0
	// grid 判空
	if grid == nil || len(grid) == 0 {
		return ans
	}
	row := len(grid)
	col := len(grid[0])
	waters := 0

	uf := NewUnionFind(grid)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '0' {
				waters++
			} else {
				direction := [][]int{
					{0, 1},  // 右
					{0, -1}, // 左
					{1, 0},  //下
					{-1, 0}, // 上
				}

				for _, dir := range direction {
					newRow := i + dir[0]
					newCol := j + dir[1]
					// 合法判断
					if newRow >= 0 && newRow < row && newCol >= 0 && newCol < col && grid[newRow][newCol] == '1' {
						uf.union(newRow*col+newCol, i*col+j)
					}
				}
			}
		}
	}
	ans = uf.count - waters
	return ans
}

type UnionFind struct {
	root  []int
	count int
}

func NewUnionFind(grid [][]byte) *UnionFind {
	row := len(grid)

	col := len(grid[0])

	count := row * col
	root := make([]int, count)
	for i := range root {
		root[i] = i
	}
	return &UnionFind{
		root:  root,
		count: count,
	}
}

// 函数传入一个 x， 返回 x的祖先
func (this *UnionFind) find(x int) int {
	// 索引和 元素相同时
	if x == this.root[x] {
		// 代表 x就是祖先
		return x
	} else {
		// 如果 索引和元素不相同，就以 元素为索引往前找
		this.root[x] = this.find(this.root[x])
	}

	return this.root[x]
}

// 合并操作：将元素x和元素y所在的集合合并
func (this *UnionFind) union(x, y int) {
	rootX := this.find(x)
	rootY := this.find(y)

	// 如果祖先不相等，执行合并
	if rootX != rootY {
		this.root[rootX] = rootY // 将一个集合的根 指向另一个集合的根
		this.count--
	}
}

type UnionFind2 struct {
	root  []int
	count int
}

func NewUnionFind2(grid [][]byte) *UnionFind2 {
	row := len(grid)
	col := len(grid[0])

	count := row * col

	root := make([]int, count)
	for i := range root {
		root[i] = i
	}

	return &UnionFind2{
		root:  root,
		count: count,
	}
}

func (this *UnionFind2) find(index int) int {
	if index == this.root[index] {
		return index
	} else {
		// 否则，以当前元素为索引来找，对应对应索引的元素值
		this.root[index] = this.find(this.root[index])
	}
	return this.root[index]
}

func (this *UnionFind2) union(x, y int) {
	rootX := this.find(x)
	rootY := this.find(y)

	if rootX != rootY {
		// 合并
		this.root[rootX] = rootY
		this.count--
	}
}

func numIslands2(grid [][]byte) int {
	ans := 0
	if len(grid) == 0 {
		return ans
	}
	// 不为空情况
	row := len(grid)
	col := len(grid[0])
	waters := 0

	uf := NewUnionFind2(grid)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '0' {
				waters++
			} else {
				directions := [][]int{
					{0, -1}, // 上
					{0, 1},  //下
					{-1, 0}, // 左
					{1, 0},  //右
				}
				for _, dir := range directions {
					NewRow := i + dir[0]
					NewCol := j + dir[1]
					if NewRow >= 0 && NewRow < row && NewCol >= 0 && NewCol < col && grid[NewRow][NewCol] == '1' {
						uf.union(NewRow*col+NewCol, i*col+j)
					}
				}
			}
		}
	}
	ans = uf.count - waters
	return ans
}
