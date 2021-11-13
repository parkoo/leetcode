package main

// 递归归并法
// 时间复杂度：O(nlgn)  空间复杂度：O(lgn)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func sortList_2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 快慢指针 将链表分成两部分
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	l1, l2 := head, slow.Next
	slow.Next = nil // 需将链表切断为两部分

	// 递归merge
	sortedL1 := sortList_2(l1)
	sortedL2 := sortList_2(l2)
	return merge(sortedL1, sortedL2)
}

// 对两条链表进行merge
func merge(l1, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	for l1 != nil {
		cur.Next = l1
		l1 = l1.Next
		cur = cur.Next
	}
	for l2 != nil {
		cur.Next = l2
		l2 = l2.Next
		cur = cur.Next
	}

	return dummyHead.Next
}
