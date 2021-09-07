package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度：O(n)  空间复杂度：O(h)  n为节点个数，h为树的深度(递归次数)
func maxDepth_1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ldepth := maxDepth_1(root.Left)
	rdepth := maxDepth_1(root.Right)

	if ldepth > rdepth {
		return ldepth + 1
	}
	return rdepth + 1
}
