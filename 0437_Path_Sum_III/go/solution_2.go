package main

// 思路: 前缀和求区间和 + DFS + 回溯
// 时间复杂度：O(n)  空间复杂度：O(n)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum_2(root *TreeNode, targetSum int) int {
	res := 0

	// prefixSumAndPathSum 记录“从根节点root出发到当前节点”的前缀和以及对应的路径数信息
	// prefixSumAndPathSum[i] 表示“从根节点root出发到当前节点”的前缀和为i的路径数
	// 以 'prefixSum' 表示从根节点root当前节点的前缀和，则 ‘prefixSum - targetSum’ 表示从根节点root到某些(/或某个节点), 其前缀和为'prefixSum - targetSum',定义这些节点为 nodes = [node_A, Node_B...]
	// 则 prefixSumAndPathSum[prefixSum - targetSum] 表示这些节点的个数, 即 len(nodes)
	// 从根节点root到当前节点的前缀和为prefixSum, 到nodes中的节点的前缀和为‘prefixSum - targetSum’,
	// 则 从nodes中的任一节点到当前节点的前缀和都为targetSum, 即以nodes中任一节点向下出发，以当前节点为终点的路径和为targetSum
	prefixSumAndPathSum := make(map[int]int)
	prefixSumAndPathSum[0] = 1 // prefixSum == targetSum 的情况

	// dfs递归遍历二叉树
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, prefixSum int) {
		if node == nil {
			return
		}

		// 计算当前节点的前缀和
		prefixSum += node.Val

		// 注意，先计算更新结果，后对当前节点的前缀和路径数进行记录, 避免targetSum == 0时，直接将当前节点算进去
		res += prefixSumAndPathSum[prefixSum-targetSum]
		prefixSumAndPathSum[prefixSum] += 1

		dfs(node.Left, prefixSum)
		dfs(node.Right, prefixSum)

		// 回溯处理，保证左右子树的 prefixSumAndPathSum 数据不会互相影响（题目定义path不跨根节点）
		prefixSumAndPathSum[prefixSum] -= 1
	}

	dfs(root, 0)

	return res
}
