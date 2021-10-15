package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历BST(二分搜索树)思路
// 时间复杂度：O(n)  空间复杂度：O(n)

func convertBST_1(root *TreeNode) *TreeNode {
	r := root
	prev := 0 // 记录前驱节点的值
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Right
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			node.Val += prev
			prev = node.Val

			root = node.Left
		}
	}

	return r
}
