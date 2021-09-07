package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度：O(n)  空间复杂度：O(h)  n为节点个数，h为树的深度
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	l := invertTree(root.Left)
	r := invertTree(root.Right)
	root.Right = l
	root.Left = r

	return root
}
