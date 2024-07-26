package main

// 思路: 递归 适合计算任何树， 没有用到完全二叉树的性质 复杂度较高

// 时间复杂度: O(n)    空间复杂度: O(n)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes_1(root *TreeNode) int {
	return helper(root)
}

func helper(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return helper(node.Left) + helper(node.Right) + 1
}
