package main

// 思路: 递归中序遍历BST, 遍历时记录已遍历节点的个数，剪枝处理

// 时间复杂度: O(N)    空间复杂度: O(N)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int = 0 // 最终结果
var cnt int = 0 // 记录已经访问的节点的个数

func kthSmallest_1(root *TreeNode, k int) int {
	res, cnt = 0, 0
	helper(root, k)
	return res
}

// 中序递归遍历BST
func helper(node *TreeNode, k int) {
	if node == nil {
		return
	}

	// 左子树
	helper(node.Left, k)

	// 当前节点
	cnt++
	if cnt == k {
		res = node.Val
		// 得到结果后直接返回，剪枝处理
		return
	}

	// 右子树
	helper(node.Right, k)
}
