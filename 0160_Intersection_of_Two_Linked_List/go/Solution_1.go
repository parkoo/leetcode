package main

// 思路：双指针

// 时间复杂度：O(m+n)  空间复杂度：O(1)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB
	for l1 != l2 { // 比较的是指针
		// 错误写法：
		// l1 = l1.Next
		// if l1 == nil {
		// 	l1 = headB
		// }
		// 注意这里必须要允许l1 、l2 为空， 因为两个链表不相交时需要最后两个指针都为空
		if l1 == nil {
			l1 = headB
		} else {
			l1 = l1.Next
		}

		// 注意这里必须要允许l1 、l2 为空， 因为两个链表不相交时需要最后两个指针都为空
		if l2 == nil {
			l2 = headA
		} else {
			l2 = l2.Next
		}
	}

	return l1
}
