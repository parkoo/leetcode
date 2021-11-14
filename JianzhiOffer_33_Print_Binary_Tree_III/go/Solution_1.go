package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 利用一个队列做层序遍历  单层结果使用双端队列保存
// 时间复杂度：O(n)  空间复杂度：O(n)

func levelOrder_1(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		// 需记录此时队列的长度
		cnt := len(queue)
		// 双端队列
		arr := make([]int, len(queue))
		for i := 0; i < cnt; i++ {
			node := queue[0]
			queue = queue[1:]
			if len(res)%2 == 0 { // 偶数层
				arr[i] = node.Val
			} else { // 奇数层
				arr[cnt-1-i] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
		res = append(res, arr)
	}

	return res
}
