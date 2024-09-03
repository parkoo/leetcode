package main

// 思路:  迭代法

// 时间复杂度: O(n)    空间复杂度: O(1)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	cur := root
	pre := root
	for cur != nil {
		pre = cur
		if val > cur.Val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}

	if val < pre.Val {
		pre.Left = &TreeNode{Val: val}
	} else {
		pre.Right = &TreeNode{Val: val}
	}

	return root
}
