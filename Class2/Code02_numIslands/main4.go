package main

type UnionFind4 struct {
	// root 存所有节点
	root []int

	// 数量,一开始每个节点都是独立的集合
	count int
}

func NewUnionFind4(grid [][]byte) *UnionFind4 {
	row, col := len(grid), len(grid[0])
	count := row * col
	root := make([]int, count)
	for i := range root {
		root[i] = i
	}
	return &UnionFind4{
		root:  root,
		count: count,
	}
}

// 传入一个 x，找出x的祖先
func (this *UnionFind4) find(x int) int {
	if x == this.root[x] {
		return x
	} else {
		this.root[x] = this.find(this.root[x])
	}
	return this.root[x]
}

// 合并, 传入两个参数，将他们合并。
/*
怎么合并？
	祖先相等不需要合并，祖先不相等，把祖先等于另一个参数的祖先
*/
func (this *UnionFind4) union(x, y int) {
	rootX := this.find(x)
	rootY := this.find(y)

	if rootX != rootY {
		//执行合并
		this.root[rootX] = rootY
		// 合并了，集合种类数量少了
		this.count--
	}
}

func numIslands5(grid [][]byte) int {
	ans := 0
	if len(grid) == 0 {
		return ans
	}
	row, col := len(grid), len(grid[0])
	wt := 0
	uf := NewUnionFind4(grid)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '0' {
				wt++
			} else {
				directions := [][]int{
					{0, -1},
					{-1, 0},
				}
				for _, val := range directions {
					NewRow := i + val[0]
					NewCol := j + val[1]
					if NewRow >= 0 && NewRow < row && NewCol >= 0 && NewCol < col && grid[i][j] == '1' {
						uf.union(NewRow*col+NewCol, i*col+j)
					}
				}
			}
		}
	}
	ans = uf.count - wt
	return ans
}
