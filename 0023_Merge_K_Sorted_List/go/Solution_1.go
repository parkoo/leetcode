package main

// 数组的分治思想 merge
// 时间复杂度：O(kn*lgk)  空间复杂度：O(lgk) k为链表个数, n为链表平均长度

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return mergeLists(lists, 0, len(lists)-1)
}

func mergeLists(lists []*ListNode, l, r int) *ListNode {
	if l >= r {
		return lists[l]
	}

	mid := l + (r-l)/2
	l1 := mergeLists(lists, l, mid)
	l2 := mergeLists(lists, mid+1, r)
	return merge(l1, l2)
}

func merge(l, r *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	for l != nil && r != nil {
		if l.Val < r.Val {
			cur.Next = l
			l = l.Next
		} else {
			cur.Next = r
			r = r.Next
		}

		cur = cur.Next
	}

	if l != nil {
		cur.Next = l
	}

	if r != nil {
		cur.Next = r
	}

	return dummyHead.Next
}
