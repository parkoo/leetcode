package main

import "math"

// 思路：两次递归
// 当前节点左子树的深度与右子树的深度之和即为以当前节点为根的二叉树的最长路径
// 两次递归中序遍历嵌套，一次为求以每个节点为根的数的深度，一次为求以每个节点为根的最长路径

// 时间复杂度：O(n^2)  空间复杂度：O(h^2)  h为树的高度，每一层递归都需要分配栈空间，每一层中分配的空间为常数

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max_path int // 通过全局变量记录最长路径

func diameterOfBinaryTree_1(root *TreeNode) int {
	max_path = 0 // 刷新全局变量

	getDiameterOfBinaryTree(root)
	return max_path
}

// 递归中序遍历每个节点，计算以该节点为根节点的最长路径，比较所有节点的最长路径获取整棵树的最长路径
func getDiameterOfBinaryTree(root *TreeNode) {
	if root == nil {
		return
	}

	// 当前节点左子树的深度与右子树的深度之和即为以当前节点为根的二叉树的最长路径
	path := getDepthOfBinaryTree(root.Left) + getDepthOfBinaryTree(root.Right)
	if path > max_path {
		max_path = path
	}

	getDiameterOfBinaryTree(root.Left)
	getDiameterOfBinaryTree(root.Right)
}

// 计算以当前节点为根节点的树的深度
func getDepthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return int(math.Max(float64(getDepthOfBinaryTree(root.Left)),
		float64(getDepthOfBinaryTree(root.Right)))) + 1
}
