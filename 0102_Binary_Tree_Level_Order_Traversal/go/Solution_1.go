package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 借助队列实现层序遍历
// 时间复杂度：O(n)  空间复杂度：O(n)

func levelOrder_1(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	cnt := 1 // 记录每层中的节点个数
	for len(queue) > 0 {
		temp := 0
		arr := make([]int, 0)
		for i := 0; i < cnt; i++ {
			node := queue[0]
			queue = queue[1:]
			arr = append(arr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
				temp++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				temp++
			}
		}
		res = append(res, arr)
		cnt = temp
	}

	return res
}
