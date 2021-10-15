package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
// 时间复杂度：O(n)  空间复杂度：O(n)

func isValidBST_1(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, minVal, maxVal int) bool {
	if root == nil {
		return true
	}

	if root.Val <= minVal || root.Val >= maxVal {
		return false
	}

	return helper(root.Left, minVal, root.Val) && helper(root.Right, root.Val, maxVal)
}
