package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度：O(n)  空间复杂度：O(h)  n为节点个数，h为树的深度
// 注意深度是指根结点到叶子节点的路径，叶子节点是指左右子节点都为nil的节点，而非nil节点本身, 需注意只有单边子节点的特殊case
func minDepth_1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 叶子节点
	if root.Left == nil && root.Right == nil {
		return 1
	}

	var mindepth int = math.MaxInt32
	if root.Left != nil {
		ldepth := minDepth_1(root.Left)
		if ldepth < mindepth {
			mindepth = ldepth
		}
	}
	if root.Right != nil {
		rdepth := minDepth_1(root.Right)
		if rdepth < mindepth {
			mindepth = rdepth
		}
	}

	return mindepth + 1
}
