package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归 (递归函数单独在外层定义，需要返回结果)， 中序遍历二叉树
// 时间复杂度：O(n)  空间复杂度：O(n)

func inorderTraversal_2(root *TreeNode) []int {
	return inorder(root)
}

// 递归函数
func inorder(root *TreeNode) []int {
	res := make([]int, 0)

	if root == nil {
		return res
	}

	res = append(res, inorder(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorder(root.Right)...)

	return res
}
