package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 转化为两棵树是否为镜像对称
// 时间复杂度：O（n） 空间复杂度：O(h)  n为root的节点个数，h为root的深度
func isSymmetric_1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	return p.Val == q.Val && symmetric(p.Left, q.Right) && symmetric(p.Right, q.Left)
}
