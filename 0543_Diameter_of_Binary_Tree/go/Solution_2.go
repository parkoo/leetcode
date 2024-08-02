package main

// 思路：一次递归遍历 递归中维护两个值  思路与lc124相似
// 当前节点左子树的深度与右子树的深度之和即为以当前节点为根的二叉树的最长路径
// 比较Solution_1，考虑在求树的深度的同时将最长路径一并求得，则只需一次递归遍历即可

// 时间复杂度：O(n)  空间复杂度：O(h)  h为树的高度，每一层递归都需要分配栈空间，每一层中分配的空间为常数

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree_2(root *TreeNode) int {
	res := 0

	// 返回当前节点的深度(非nil的根节点的深度为1)
	// 在递归过程中统计以当前节点为根节点的最大直径
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftDepth, rightDepth := dfs(node.Left), dfs(node.Right)

		// 在递归过程中统计以当前节点为根节点的最大直径
		// 以当前节点为根节点的最大直径等于当前节点左子树的深度+当前节点右子树的深度
		curPathLen := leftDepth + rightDepth
		if curPathLen > res {
			res = curPathLen
		}

		// 返回当前节点的深度
		return max(leftDepth, rightDepth) + 1
	}

	dfs(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
