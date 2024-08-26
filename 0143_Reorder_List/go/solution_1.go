package main

// 思路: 快慢指针切分链表 + 反转链表 + 合并链表

// 时间复杂度: O(n)    空间复杂度: O(1)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList_1(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 快慢指针切分链表
	s, f := head, head.Next
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	l1, l2 := head, s.Next
	s.Next = nil

	// 反转链表
	var pre *ListNode
	cur := l2
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	l2 = pre

	// 合并链表
	for l2 != nil {
		next1 := l1.Next
		next2 := l2.Next
		l1.Next = l2
		l2.Next = next1
		l2 = next2
		l1 = l1.Next.Next
	}
}
