package main

// 链表排序
// 递归归并法
// 时间复杂度：O(nlgn)  空间复杂度：O(lgn)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeSortList_1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 快慢指针 将链表均分成两部分
	// 这里fast必须取head.Next, 若取了head, 则在链表节点为偶数时，切分不均匀，导致时间复杂度和空间复杂度上升
	// 或者这里先加个虚拟头结点dummyHead, 可以令slow, fast := dummyHead, dummyHead, 效果也是均分的
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	l1, l2 := head, slow.Next
	slow.Next = nil // 需将链表切断为两部分

	// 递归merge
	sortedL1 := mergeSortList_1(l1)
	sortedL2 := mergeSortList_1(l2)
	return mergeList_1(sortedL1, sortedL2)
}

// 对两条链表进行merge 需要返回值
func mergeList_1(l1, l2 *ListNode) *ListNode {
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
