package main

// 思路：寻找前驱节点

// 时间复杂度：O(n)  空间复杂度：O(1)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten_2(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left == nil {
			cur = cur.Right
			continue
		}

		// 寻找当前节点右子树的前驱结点
		// 该前驱结点为当前节点左子树的最右侧节点
		pre := cur.Left
		for pre.Right != nil {
			pre = pre.Right
		}

		// 更新当前节点右子树的前驱节点
		pre.Right = cur.Right
		// 更新当前节点
		cur.Right = cur.Left
		cur.Left = nil

		cur = cur.Right
	}
}
