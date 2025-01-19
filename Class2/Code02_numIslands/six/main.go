package six

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	uf := NewUnionFind(grid)
	wt := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			// 如果是水
			if grid[i][j] == '0' {
				wt++
			} else {
				// 陆地
				directions := [][]int{
					{0, 1},
					{1, 0},
				}
				for _, v := range directions {
					NewRow := i + v[0]
					NewCol := j + v[1]
					if NewRow >= 0 && NewRow < row && NewCol >= 0 && NewCol < col && grid[NewRow][NewCol] == '1' {
						// 合并
						uf.union(NewRow*col+NewCol, i*col+j)
					}
				}
			}
		}
	}
	ans := uf.count - wt
	return ans
}

// 并查集
type UnionFind struct {
	root  []int
	count int
}

func NewUnionFind(grid [][]byte) *UnionFind {
	row, col := len(grid), len(grid[0])
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

// find 找x的祖先
func (this *UnionFind) find(x int) int {
	if x == this.root[x] {
		return x
	} else {
		this.root[x] = this.find(this.root[x])
	}
	return this.root[x]
}

func (this *UnionFind) union(x, y int) {
	rootX := this.find(x)
	rootY := this.find(y)
	if rootX != rootY {
		// 执行合并
		this.root[rootX] = rootY
		this.count--
	}
}
