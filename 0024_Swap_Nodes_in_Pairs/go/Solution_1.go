package main

// 虚拟头节点
// 时间复杂度：O(n)  空间复杂度：O(1)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := new(ListNode)
	dummyHead.Next = head
	pre := dummyHead

	slow, fast := head, head.Next
	for fast != nil {
		next := fast.Next
		fast.Next = slow
		slow.Next = next
		pre.Next = fast
		pre = slow
		slow = slow.Next
		if slow != nil {
			fast = slow.Next
		} else {
			fast = nil
		}
	}

	return dummyHead.Next
}
