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

var maxNodeNum int = 0

func diameterOfBinaryTree_2(root *TreeNode) int {
	maxNodeNum = 0
	dfs(root)

	res := maxNodeNum - 1 // 节点个数比路径个数多1
	return res
}

// 遍历统计以每个节点为根节点的最大节点个数
func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftNum := dfs(node.Left)
	rightNum := dfs(node.Right)
	if leftNum+rightNum+1 > maxNodeNum {
		maxNodeNum = leftNum + rightNum + 1
	}

	theMax := leftNum
	if rightNum > leftNum {
		theMax = rightNum
	}

	return theMax + 1

}
