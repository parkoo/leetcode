package main

// 思路: 二叉树“之”字形遍历

// 时间复杂度: O(n)    空间复杂度: O(n)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)
	level := 0

	for len(q) > 0 {
		cnt := len(q)
		item := make([]int, 0)
		for i := 0; i < cnt; i++ {
			cur := q[0]
			q = q[1:]
			item = append(item, cur.Val)

			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}

		// 奇数层需要反转结果
		if level%2 == 1 {
			i, j := 0, len(item)-1
			for i < j {
				item[i], item[j] = item[j], item[i]
				i++
				j--
			}
		}
		res = append(res, item)

		level++
	}

	return res
}
