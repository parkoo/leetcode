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

var maxSum int = math.MinInt16

func maxPathSum_1(root *TreeNode) int {
	maxSum = math.MinInt16
	helper(root)
	return maxSum
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 直接过滤掉小于0的sum
	leftSum := max(helper(root.Left), 0)
	rightSum := max(helper(root.Right), 0)

	// 更新全局最大值
	curMaxSum := root.Val + leftSum + rightSum
	if curMaxSum > maxSum {
		maxSum = curMaxSum
	}

	// 计算递归返回值，左右子树只可取其一，以保证路径不分叉
	sum := root.Val + max(leftSum, rightSum)
	return sum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
