package main

// 双指针 快慢指针  Floyd判环算法（龟兔赛跑算法）
// 时间复杂度：O(n)  空间复杂度：O(1)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle_3(head *ListNode) bool {

	// 以 fast != nil && fast.Next != nil 作为循环判断
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}
