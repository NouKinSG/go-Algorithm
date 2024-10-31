package five

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}
	wt := 0
	uf := NewUnionFind(grid)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 判断是岛屿还是海
			if grid[i][j] == '0' {
				// 海水
				wt++
			} else {
				// 陆地
				diraction := [][]int{
					{0, -1},
					{-1, 0},
				}
				for _, val := range diraction {
					// 新行、新列
					newRow := i + val[0]
					newCol := j + val[1]
					if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && grid[newRow][newCol] == '1' {
						// 新的和旧的合并
						uf.Uninx(newRow*n+newCol, i*n+j)
					}
				}
			}
		}
	}
	ans := uf.count - wt
	return ans
}

type UnionFind struct {
	root  []int
	count int
}

// 构造
func NewUnionFind(grid [][]byte) *UnionFind {
	m, n := len(grid), len(grid[0])
	root := make([]int, m*n)
	for i := range root {
		root[i] = i
	}

	count := m * n
	return &UnionFind{
		root:  root,
		count: count,
	}
}

// find
func (this *UnionFind) find(x int) int {
	if x == this.root[x] {
		return x
	} else {
		this.root[x] = this.find(this.root[x])
	}

	return this.root[x]
}

// Uninx
func (this *UnionFind) Uninx(x, y int) {
	rootX := this.find(x)
	rootY := this.find(y)
	if rootX != rootY {
		this.root[rootX] = rootY
		this.count--
	}
}
