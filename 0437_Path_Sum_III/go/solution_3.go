package main

// 思路：递归法 同solution_1类似

// 时间复杂度：O(n^2)  空间复杂度：O(n)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var cnt int = 0

func pathSum_3(root *TreeNode, targetSum int) int {
	cnt = 0
	dfs(root, targetSum)
	return cnt
}

func dfs(node *TreeNode, target int) {
	if node == nil {
		return
	}

	helper(node, target)
	dfs(node.Left, target)
	dfs(node.Right, target)
}

func helper(node *TreeNode, target int) {
	if node == nil {
		return
	}

	if node.Val == target {
		cnt++
	}

	helper(node.Left, target-node.Val)
	helper(node.Right, target-node.Val)
}
