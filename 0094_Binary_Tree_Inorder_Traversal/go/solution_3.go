package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归 (递归函数在内层定义，可共享外层函数的局部结果变量res)， 中序遍历二叉树
// 时间复杂度：O(n)  空间复杂度：O(n)

func inorderTraversal_3(root *TreeNode) []int {
	res := make([]int, 0)

	// 声明并定义递归函数
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}

	// 调用递归函数
	inorder(root)

	return res
}
