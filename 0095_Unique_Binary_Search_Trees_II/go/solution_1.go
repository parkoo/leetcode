package main

// 思路:  递归法

// 时间复杂度: O(?)    空间复杂度: O(?)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {

	return dfs(1, n)
}

// 由[start, end]节点可以构成的不同单位二分搜索树
func dfs(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	curTrees := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		leftTrees := dfs(start, i-1)
		rightTrees := dfs(i+1, end)

		for _, leftTree := range leftTrees {
			for _, rightTree := range rightTrees {

				// 注意：每选定一对 leftTree 和 rightTree 后，都必须生成一个新的root
				curRoot := &TreeNode{
					Val: i,
				}
				curRoot.Left = leftTree
				curRoot.Right = rightTree
				curTrees = append(curTrees, curRoot)
			}
		}
	}

	return curTrees
}
