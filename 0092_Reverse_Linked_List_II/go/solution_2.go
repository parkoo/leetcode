package main

// 思路：截断区间 反转链表

// 时间复杂度：O(n)  空间复杂度：O(1)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween_2(head *ListNode, left int, right int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head

	pre, end := dummyHead, dummyHead
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	for i := 0; i < right; i++ {
		end = end.Next
	}

	start := pre.Next
	next := end.Next
	curStart, curEnd := reverse(start, end)
	pre.Next = curStart
	curEnd.Next = next

	return dummyHead.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode
	cur := head
	tail.Next = nil

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return tail, head
}
