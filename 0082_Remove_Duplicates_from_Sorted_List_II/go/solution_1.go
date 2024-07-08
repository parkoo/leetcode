package main

// 思路: 双指针

// 时间复杂度: O(n)    空间复杂度: O(1)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates_1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := new(ListNode)
	dummyHead.Next = head

	pre := dummyHead
	left, right := head, head.Next
	for right != nil {
		if right.Val != left.Val {
			right = right.Next
			left = left.Next
			pre = pre.Next
			continue
		}

		for right.Next != nil && right.Next.Val == left.Val {
			right = right.Next
		}

		next := right.Next
		pre.Next = next
		left = next
		right = next
		if next != nil {
			right = right.Next
		}
	}

	return dummyHead.Next
}
