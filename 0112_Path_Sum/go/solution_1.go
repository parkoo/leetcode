package main

// 思路: 递归

// 时间复杂度: O(n)    空间复杂度: O(h) h为树的高度

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum_1(root *TreeNode, targetSum int) bool {
	// 需要单独处理root == nil的情况
	if root == nil {
		return false
	}

	return helper(root, targetSum)
}

func helper(node *TreeNode, target int) bool {
	// 这里是到达非nil的叶子节点
	if node.Left == nil && node.Right == nil {
		return target == node.Val
	}

	nextTarget := target - node.Val
	res := false
	if node.Left != nil {
		res = res || helper(node.Left, nextTarget)
	}
	if node.Right != nil {
		res = res || helper(node.Right, nextTarget)
	}

	return res
}
