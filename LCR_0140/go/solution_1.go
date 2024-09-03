package main

// 思路:  双指针

// 时间复杂度: O(n)    空间复杂度: O(1)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func trainingPlan(head *ListNode, cnt int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head

	s, f := dummyHead, dummyHead
	for i := 0; i < cnt; i++ {
		f = f.Next
	}

	for f != nil {
		s = s.Next
		f = f.Next
	}

	return s
}
