package main

type UnionFind3 struct {
	root  []int
	count int
}

func NewUnionFind3(grid [][]byte) *UnionFind3 {
	row := len(grid)
	col := len(grid[0])
	count := row * col
	root := make([]int, count)

	// 赋初始值
	for i := range root {
		root[i] = i
	}
	return &UnionFind3{
		root:  root,
		count: count,
	}
}

func (this *UnionFind3) find(x int) int {
	if x == this.root[x] {
		return x
	} else {
		this.root[x] = this.find(this.root[x])
	}
	return this.root[x]
}

func (this *UnionFind3) union(x, y int) {
	// 合并前先找祖先
	rootX := this.find(x)
	rootY := this.find(y)
	if rootX != rootY {
		this.root[rootX] = rootY
		this.count--
	}
}

func numIslands4(grid [][]byte) int {
	ans := 0
	if len(grid) == 0 {
		return ans
	}
	row, col := len(grid), len(grid[0])
	wt := 0
	uf := NewUnionFind3(grid)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '0' {
				wt++
			} else {
				directions := [][]int{
					{0, -1}, // 上
					{-1, 0}, // 左
				}
				for _, val := range directions {
					NewRow := i + val[0]
					NewCol := j + val[1]
					if NewRow >= 0 && NewRow < row && NewCol >= 0 && NewCol < col && grid[NewRow][NewCol] == '1' {
						uf.union(NewRow*col+NewCol, i*col+j)
					}
				}
			}
		}
	}
	ans = uf.count - wt
	return ans
}
