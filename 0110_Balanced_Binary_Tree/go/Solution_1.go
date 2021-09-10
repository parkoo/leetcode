package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度：O（n） 空间复杂度：O(h)  n为root的节点个数，h为root的深度
func isBalanced(root *TreeNode) bool {
	return depth(root) != -1
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lDepth := depth(root.Left)
	rDepth := depth(root.Right)

	if lDepth == -1 || rDepth == -1 || abs(lDepth, rDepth) > 1 {
		return -1
	}

	return max(lDepth, rDepth) + 1
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
