package main

type TreeNode struct {
	val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func findLCA(node1, node2 *TreeNode) *TreeNode {
	p, q := node1, node2
	for p != q {
		if p.Parent == nil {
			p = node2
		} else {
			p = p.Parent
		}

		if q.Parent == nil {
			q = node1
		} else {
			q = q.Parent
		}
	}
	return p
}
