package main

// map 记录
// 时间复杂度：O(n)  空间复杂度：O(n)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle_1(head *ListNode) bool {
	seen := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = struct{}{}
		head = head.Next
	}

	return false
}
