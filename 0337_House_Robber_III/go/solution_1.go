package main

// 思路: 树形动态规划

// 时间复杂度: O(n)    空间复杂度: O(n)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	notRobVal, robVal := dfs(root)
	return max(notRobVal, robVal)
}

// 计算遍历到当前节点时, 选择"偷取该节点"以及"不偷取该节点"可获得的全局最大值
// 返回两个值, 分别表示选择"偷取该节点"以及"不偷取该节点"可获得的全局最大值
func dfs(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}

	lNotRob, lRob := dfs(node.Left)
	rNotRob, rRob := dfs(node.Right)

	// 若当前选择“不偷取”, 则左右子节点均可选择“偷取”或者“不偷取”，从中选取较大值
	curNotRob := max(lNotRob, lRob) + max(rNotRob, rRob)

	// 若当前选择“偷取”, 则左右子节点均只可选择“不偷取”
	curRob := node.Val + lNotRob + rNotRob

	return curNotRob, curRob
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
