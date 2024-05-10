package main

// 思路: 层序遍历（广度优先遍历）二叉树, 借助一个队列实现

// 时间复杂度: O(n)    空间复杂度: O(n)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView_1(root *TreeNode) []int {
	res := make([]int, 0)

	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		// 计算当前层节点的个数
		cnt := len(queue)

		// 遍历当前层
		for i := 0; i < cnt; i++ {
			// 出队列
			node := queue[0]
			queue = queue[1:]
			if i == cnt-1 {
				res = append(res, node.Val)
			}

			// 左右节点入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return res
}
