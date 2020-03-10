package main

import "math"

// 当前节点左子树的深度与右子树的深度之和即为以当前节点为根的二叉树的最长路径
// 比较Solution_1，考虑在求树的深度的同时将最长路径一并求得，则只需一次递归中序遍历即可
// 时间复杂度：O(n)  空间复杂度：O(h)  h为树的高度，每一层递归都需要分配栈空间，每一层中分配的空间为常数
 
/* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/

var max_path_2 int

func diameterOfBinaryTree_2(root *TreeNode) int {
	max_path_2 = 0  // 刷新全局变量

	getDepthAndDiameterOfBinaryTree(root)
	return max_path_2		
}

// 计算以当前节点为根节点的树的深度
func getDepthAndDiameterOfBinaryTree(root *TreeNode) int {
	if root==nil {
		return 0
	}

	depth_left := getDepthAndDiameterOfBinaryTree(root.Left)
	depth_right:= getDepthAndDiameterOfBinaryTree(root.Right)

	// 当前节点左子树的深度与右子树的深度之和即为以当前节点为根的二叉树的最长路径
	path := depth_left + depth_right
	if path>max_path_2 {
		max_path_2 = path
	}

	return int(math.Max(float64(depth_left), float64(depth_right))) + 1
}