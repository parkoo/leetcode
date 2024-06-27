package main

// 双指针 快慢指针  Floyd判环算法（龟兔赛跑算法）
// 时间复杂度：O(n)  空间复杂度：O(1)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle_2(head *ListNode) bool {
	slow, fast := head, head
	if fast == nil || fast.Next == nil {
		return false
	}
	slow = slow.Next
	fast = fast.Next.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
