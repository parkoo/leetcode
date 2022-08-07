package main

// 单调栈

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	res := make([]int, 0)

	stack := make([]int, 0) // 单调栈，存下标
	cur := head
	i := 0 // 下标
	for cur != nil {
		for len(stack) > 0 && cur.Val > res[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = cur.Val
			stack = stack[:len(stack)-1]
		}

		res = append(res, cur.Val)
		stack = append(stack, i)
		i++
		cur = cur.Next
	}

	for len(stack) > 0 {
		res[stack[len(stack)-1]] = 0
		stack = stack[:len(stack)-1]
	}

	return res
}
