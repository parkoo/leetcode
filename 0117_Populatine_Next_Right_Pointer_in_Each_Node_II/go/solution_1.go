package main

// 思路:  二叉树的层序遍历

// 时间复杂度: O(n)    空间复杂度: O(n)

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect_1(root *Node) *Node {
	if root == nil {
		return root
	}

	q := make([]*Node, 0)
	q = append(q, root)

	for len(q) > 0 {

		cnt := len(q)
		for i := 0; i < cnt; i++ {
			cur := q[0]
			q = q[1:]

			if i < cnt-1 {
				cur.Next = q[0]
			}

			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}

	return root
}
