package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 非递归前序遍历 借助一个栈
// 时间复杂度：O(n)  空间复杂度：O(n)

func preorderTraversal_1(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return res
}
