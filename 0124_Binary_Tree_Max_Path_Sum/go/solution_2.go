package main

import "math"

// 思路：递归法 两次递归 思路与lc543相似
// 两次递归可以合并在一个递归中处理，见solution_1

// 时间复杂度：O(n^2)  空间复杂度：O(n^2)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int = math.MinInt

func maxPathSum_2(root *TreeNode) int {
	res = math.MinInt

	dfs(root)

	return res
}

// 递归遍历树的每个节点，遍历到当前节点时，计算以当前节点为根节点的最大路径和
// 以遍历到的节点为根节点，可以包含左子树或右子树或左右子树或仅仅只是该根节点，主要看哪种情况计算的路径和的值最大
func dfs(node *TreeNode) {
	if node == nil {
		return
	}

	// 递归主体处理逻辑
	leftSum := sum(node.Left)
	rightSum := sum(node.Right)

	// 选择加上值大于0的左子树和值大于0的右子树
	curMax := node.Val + max(leftSum, 0) + max(rightSum, 0) // 处理节点值为负值的情况
	if curMax > res {
		res = curMax
	}

	dfs(node.Left)
	dfs(node.Right)
}

// 递归计算以当前节点为根节点（仅包含左子树或者右子树）的最大路径和
func sum(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftSum := sum(node.Left)
	rightSum := sum(node.Right)

	// 左、右子树中选择值较大的且不能为负
	maxVal := max(max(leftSum, rightSum), 0) // 处理节点为负值的情况

	curVal := maxVal + node.Val

	return curVal
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
