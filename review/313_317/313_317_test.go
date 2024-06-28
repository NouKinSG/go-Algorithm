package _13_317

import "testing"

func TestTrap(t *testing.T) {
	height := []int{4, 2, 0, 3, 2, 5}
	t.Log(trap(height), "--------------")
}

func TestPathSum(t *testing.T) {
	// 测试路径总和 III
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: -2,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			Right: &TreeNode{
				Val: -3,
				Right: &TreeNode{
					Val: 11,
				},
			},
		},
		Right: &TreeNode{
			Val: -3,
			Right: &TreeNode{
				Val: 11,
			},
		},
	}
	t.Log(pathSum(root, 8))
}

func TestNumIslands(t *testing.T) {
	// 测试NumIslands

	// 岛屿数量
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	t.Log(numIslands(grid))

}

func TestMinPathSum(t *testing.T) {
	// 测试最小路径和
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	t.Log(minPathSum(grid))
}
