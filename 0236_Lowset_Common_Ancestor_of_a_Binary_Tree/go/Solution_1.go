package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
// 若 root 为 p, q 的最近公共祖先， 则只有三种情况：
//   <1> p, q分散在root的左右子树中
//   <2> root == p, q 在 root 的左子树或右子树中
//   <3> root == q, p 在 root 的左子树或右子树中

// 时间复杂度：O(n)  空间复杂度：O(n)

func lowestCommonAncestor_1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 返回值是待寻找公共节点的两个节点之一
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	// 递归的返回值若不为空，则便已经是最近的祖先
	left := lowestCommonAncestor_1(root.Left, p, q)
	right := lowestCommonAncestor_1(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}
