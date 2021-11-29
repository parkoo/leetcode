package main

// 虚拟头节点 选取分区法
// 时间复杂度：O(n)  空间复杂度：O(1)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition_1(head *ListNode, x int) *ListNode {
	// 两个虚拟头节点 对链表进行选取分区
	small, large := new(ListNode), new(ListNode)
	curSmall, curLarge := small, large
	for head != nil {
		if head.Val < x {
			curSmall.Next = head
			curSmall = curSmall.Next
		} else {
			curLarge.Next = head
			curLarge = curLarge.Next
		}
		head = head.Next
	}

	curSmall.Next = large.Next
	curLarge.Next, large = nil, nil
	return small.Next
}
