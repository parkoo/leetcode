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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	return dfs(root, p, q)
}

// 定义递归函数dfs， 表示求解以node为根节点时，p、q的公共祖先
func dfs(node, p, q *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	// 当前节点是待寻找公共节点的两个节点之一
	if node == p || node == q {
		return node
	}

	// 递归当前节点的左右子树
	left := dfs(node.Left, p, q)
	right := dfs(node.Right, p, q)
	// 当前节点的左右子树分别包含p、q，则当前节点就是p、q的最近公共祖先
	if left != nil && right != nil {
		return node
	}

	// 当前节点的左子树中包含p、q
	if left != nil {
		return left
	}

	// 当前节点的右子树中包含p、q
	return right
}
