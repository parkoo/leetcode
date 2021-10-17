package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
// 时间复杂度：O(n^2)  空间复杂度：O(n)

func pathSum_1(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	res := getSum(root, targetSum)
	res += pathSum_1(root.Left, targetSum)
	res += pathSum_1(root.Right, targetSum)
	return res
}

// 计算以root为根节点的和为targetSum的路径总数
func getSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	sum := 0
	if root.Val == targetSum {
		sum++
	}

	targetSum -= root.Val
	sum += getSum(root.Left, targetSum)
	sum += getSum(root.Right, targetSum)

	return sum
}
