package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历 全局变量记录状态
// 时间复杂度：O(n)  空间复杂度：O(lgn)

var root *TreeNode
var Pre *TreeNode

func Convert(pRootOfTree *TreeNode) *TreeNode {
	if pRootOfTree == nil {
		return nil
	}

	Convert(pRootOfTree.Left)
	if root == nil {
		root = pRootOfTree
	}
	if Pre != nil {
		pRootOfTree.Left = Pre
		Pre.Right = pRootOfTree
	}
	Pre = pRootOfTree
	Convert(pRootOfTree.Right)

	return root
}
