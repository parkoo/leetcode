package main

// 思路:  二叉树前序遍历, 通过上层向下层传递信息

// 时间复杂度: O(n)    空间复杂度: O(n)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sum := 0
	if root == nil {
		return sum
	}

	var dfs func(node *TreeNode, pathNum int)
	dfs = func(node *TreeNode, pathNum int) {
		if node.Left == nil && node.Right == nil {
			pathNum = pathNum*10 + node.Val
			sum += pathNum
			return
		}

		pathNum = pathNum*10 + node.Val // 计算从上向下到达当前节点时，路径所代表的值
		if node.Left != nil {
			dfs(node.Left, pathNum)
		}
		if node.Right != nil {
			dfs(node.Right, pathNum)
		}
	}

	dfs(root, 0)

	return sum
}
