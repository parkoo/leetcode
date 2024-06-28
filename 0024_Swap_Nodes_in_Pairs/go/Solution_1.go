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
	pre := dummyHead // 记录待翻转区间的前驱节点

	slow, fast := head, head.Next
	for fast != nil {
		// 记录待翻转区间的后继节点
		next := fast.Next

		// 翻转待翻转区间
		fast.Next = slow

		// 连接翻转去区间的前驱结点和后继节点
		pre.Next = fast
		slow.Next = next

		// 翻转区间后移
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
