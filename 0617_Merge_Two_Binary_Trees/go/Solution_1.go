package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度：O(max(m,n))  空间复杂度：O(max(m,n))  m,n 为root1、root2的节点个数
func mergeTrees_1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	var val1, val2 int = 0, 0
	var left1, left2, right1, right2 *TreeNode = nil, nil, nil, nil
	if root1 != nil {
		val1 = root1.Val
		left1, right1 = root1.Left, root1.Right
	}
	if root2 != nil {
		val2 = root2.Val
		left2, right2 = root2.Left, root2.Right
	}

	val := val1 + val2
	root := &TreeNode{Val: val}

	root.Left = mergeTrees_1(left1, left2)
	root.Right = mergeTrees_1(right1, right2)

	return root
}
