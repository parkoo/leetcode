package main

// 思路: 层序遍历

// 时间复杂度: O(n)    空间复杂度: O(n)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)

	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		cnt := len(q)
		total := float64(0)
		for i := 0; i < cnt; i++ {
			cur := q[0]
			q = q[1:]
			total += float64(cur.Val)

			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}

		avg := total / float64(cnt)
		res = append(res, avg)
	}

	return res
}
