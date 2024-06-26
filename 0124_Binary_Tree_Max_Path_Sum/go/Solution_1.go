package main

import "math"

// 思路：递归法, 一次递归，递归过程中需要考虑两个值

// 时间复杂度：O(n)  空间复杂度：O(n)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int = math.MinInt

func maxPathSum(root *TreeNode) int {
	res = math.MinInt

	dfs(root)

	return res
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftSum := dfs(node.Left)
	rightSum := dfs(node.Right)

	// 计算以当前节点为根节点（可同时选择左、右子树）的最大路径和
	// 用于和最终结果比较
	// 该最大值的计算依赖于递归的返回值
	curMaxVal := node.Val + max(leftSum, 0) + max(rightSum, 0) // 处理节点值为负值的情况
	if curMaxVal > res {
		res = curMaxVal
	}

	// 计算以当前节点为根节点（只可选择左子树或者右子树之一）的最大路径和
	// 用于递归的返回值
	theMax := max(leftSum, rightSum)
	curVal := node.Val + max(theMax, 0) // 处理节点值为负值的情况

	return curVal
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
