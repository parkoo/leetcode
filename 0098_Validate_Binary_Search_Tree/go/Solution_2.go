package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 遍历法 二叉搜索树的中序遍历结果是递增序列
// 时间复杂度：O(n)  空间复杂度：O(n)

var preVal int = math.MinInt64

func isValidBST_2(root *TreeNode) bool {
	preVal = math.MinInt64
	return helper(root)
}

func helper(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if helper(root.Left) {
		if root.Val > preVal {
			preVal = root.Val
			return helper(root.Right)
		}
	}

	return false
}
