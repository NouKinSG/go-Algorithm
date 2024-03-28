package fd_test

import "testing"

// 200. 岛屿数量
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	ans := 0
	row, col := len(grid), len(grid[0])
	wt := 0
	uf := NewUninoFind(grid)

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
						// 合并
						uf.union(newRow*col+newCol, i*col+j)
					}
				}
			}
		}
	}
	ans = uf.count - wt
	return ans
}

// 并查集
type UnionFind struct {
	// 集合存储元素
	root []int

	// 统计个数
	count int
}

// 构造函数
func NewUninoFind(grid [][]byte) *UnionFind {
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

// find
func (this *UnionFind) find(x int) int {
	if x == this.root[x] {
		return x
	} else {
		this.root[x] = this.find(this.root[x])
	}
	return this.root[x]
}

// 合并
func (this *UnionFind) union(x, y int) {
	rootX := this.root[x]
	rootY := this.root[y]
	if rootX != rootY {
		this.root[rootX] = rootY
		this.count--
	}
}

func TestUF(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	ans := numIslands(grid)
	t.Log(ans)
}
