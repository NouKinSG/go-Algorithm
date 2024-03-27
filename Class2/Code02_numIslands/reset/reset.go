package reset

func numIsland(grid [][]byte) int {
	ans := 0
	row, col := len(grid), len(grid[0])
	count := row * col
	if len(grid) == 0 {
		return 0
	}
	wt := 0
	uf := NewUnionFind(grid)

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
					newRow := i + val[0]
					newCol := j + val[1]
					if newRow >= 0 && newRow < row && newCol >= 0 && newCol < col && grid[newRow][newCol] == '1' {
						uf.union(newRow*col+newCol, i*col+j)
					}
				}
			}
		}
	}
	ans = uf.count - wt
	return ans
}

type UnionFind struct {
	// 存储节点
	root []int
	// 存储数量
	count int
}

// 构造函数
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

// find，找x的祖先
func (this *UnionFind) find(x int) int {
	if x == this.root[x] {
		return x
	} else {
		this.root[x] = this.find(this.root[x])
	}
	return this.root[x]
}

// union 合并x，y
func (this *UnionFind) union(x, y int) {
	rootX := this.find(x)
	rootY := this.find(y)

	if rootX != rootY {
		this.root[rootX] = rootY
		this.count--
	}
}
