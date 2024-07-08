package main

// 思路：头插法 反转

// 时间复杂度：O(n)  空间复杂度：O(1)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween_1(head *ListNode, left int, right int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head

	start := dummyHead
	for i := 0; i < left-1; i++ {
		start = start.Next
	}

	cur := start.Next
	// 将cur后面的节点进行头插，头插后cur自动后移 不需要cur = cur.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = start.Next
		start.Next = next
	}

	return dummyHead.Next
}
