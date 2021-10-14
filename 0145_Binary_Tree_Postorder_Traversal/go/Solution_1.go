package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 非递归后序遍历 借助两个栈实现
// 时间复杂度：O(n)  空间复杂度：O(n)

func postorderTraversal_1(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack1, stack2 := make([]*TreeNode, 0), make([]*TreeNode, 0)
	stack1 = append(stack1, root)
	for len(stack1) > 0 {
		node := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, node)

		if node.Left != nil {
			stack1 = append(stack1, node.Left)
		}

		if node.Right != nil {
			stack1 = append(stack1, node.Right)
		}
	}

	for i := len(stack2) - 1; i >= 0; i-- {
		res = append(res, stack2[i].Val)
	}

	return res
}
