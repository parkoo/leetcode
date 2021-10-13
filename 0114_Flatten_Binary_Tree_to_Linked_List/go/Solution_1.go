package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 寻找前驱节点
// 时间复杂度：O(n)  空间复杂度：O(1)

func flatten_1(root *TreeNode) {
	if root == nil {
		return
	}

	curRoot := root
	for curRoot.Left != nil || curRoot.Right != nil {
		if curRoot.Left != nil {
			cur := curRoot.Left
			for cur.Right != nil {
				cur = cur.Right
			}
			cur.Right = curRoot.Right
			curRoot.Right = curRoot.Left
			curRoot.Left = nil
		}
		if curRoot.Right != nil {
			curRoot = curRoot.Right
		}
	}
}
