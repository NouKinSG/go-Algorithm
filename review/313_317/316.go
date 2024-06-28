package _13_317

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}
	wt := 0
	uf := NewUnionFind(grid)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				wt++
			} else {
				diraction := [][]int{
					{0, -1}, // 左
					{-1, 0}, // 上
				}
				for _, val := range diraction {
					newRow := i + val[0]
					newCol := j + val[1]
					if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && grid[newRow][newCol] == '1' {
						uf.Union(newRow*n+newCol, i*n+j)
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
	// 连通分量个数
	count int

	// 存储若干棵树
	root []int
}

// 初始化
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

// 查
func (this *UnionFind) find(x int) int {
	if x == this.root[x] {
		return x
	} else {
		this.root[x] = this.find(this.root[x])
	}
	return this.root[x]
}

// 合并
func (this *UnionFind) Union(x, y int) {
	// 找出祖先
	rootX := this.find(x)
	rootY := this.find(y)

	// 执行合并
	if rootX != rootY {
		this.root[rootX] = rootY
		this.count--
	}
}
