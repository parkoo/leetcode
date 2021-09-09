package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度：O（min(n1,n2)） 空间复杂度：O(min(h1, h2))  n1、n2分别为p、q的节点个数，h1、h2分别为p、q的深度
func isSameTree_1(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	lSame := isSameTree_1(p.Left, q.Left)
	rSame := isSameTree_1(p.Right, q.Right)

	return p.Val == q.Val && lSame && rSame
}
