package main

import "math"

// 思路: 中序遍历 同lc0783

// 时间复杂度: O(n)    空间复杂度: O(n)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	res := math.MaxInt64
	if root == nil {
		return 0
	}

	preVal := -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		if preVal != -1 && node.Val-preVal < res {
			res = node.Val - preVal
		}
		preVal = node.Val
		dfs(node.Right)
	}

	dfs(root)

	return res
}
