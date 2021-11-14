package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 利用一个队列 + 两个栈 实现
// 时间复杂度：O(n)  空间复杂度：O(n)

func levelOrder_2(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	stack1, stack2 := make([]*TreeNode, 0), make([]*TreeNode, 0)

	queue = append(queue, root)
	for len(queue) > 0 || len(stack1) > 0 || len(stack2) > 0 {
		arr := make([]int, 0)
		if len(queue) > 0 {
			// 打印队列里的节点
			for len(queue) > 0 {
				node := queue[0]
				queue = queue[1:]
				arr = append(arr, node.Val)

				// 将队列中节点的子节点加入第一个栈，实现下层的反转
				if node.Left != nil {
					stack1 = append(stack1, node.Left)
				}
				if node.Right != nil {
					stack1 = append(stack1, node.Right)
				}
			}
		} else {
			for len(stack1) > 0 {
				// 打印第一个栈内的节点
				node := stack1[len(stack1)-1]
				stack1 = stack1[:len(stack1)-1]
				arr = append(arr, node.Val)

				// 将第一个栈内节点的子节点加入第二个栈
				if node.Right != nil {
					stack2 = append(stack2, node.Right)
				}
				if node.Left != nil {
					stack2 = append(stack2, node.Left)
				}
			}

			// 将第二个栈中的节点放入队列实现反转
			for len(stack2) > 0 {
				node := stack2[len(stack2)-1]
				stack2 = stack2[:len(stack2)-1]
				queue = append(queue, node)
			}
		}

		res = append(res, arr)
	}

	return res
}
