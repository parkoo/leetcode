package main

// 迭代式归并
// 时间复杂度：O(nlgn)  空间复杂度：O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func sortList_3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 计算链表长度
	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}

	dummyHead := &ListNode{}
	dummyHead.Next = head

	// 对 (prev, prev+step]  [prev+step+1, succ) merge
	for step := 1; step < n; step = 2 * step {
		prev := dummyHead
		for prev != nil {
			// 寻找 mid、 succ
			mid := prev
			for j := 0; j <= step; j++ {
				if mid != nil {
					mid = mid.Next
				}
			}

			succ := prev
			for j := 0; j <= 2*step; j++ {
				if succ != nil {
					succ = succ.Next
				}
			}

			// 对（prev, mid） [mid, succ) 进行merge
			merge_3(prev, mid, succ)

			// 更新prev
			for j := 0; j < 2*step; j++ {
				if prev != nil {
					prev = prev.Next
				}
			}
		}
	}

	return dummyHead.Next
}

// 将 (prev, succ) 分成 (prev, mid-1]  [mid, succ) 进行归并
func merge_3(prev, mid, succ *ListNode) {
	cur := prev
	l1, l2 := prev.Next, mid
	for l1 != mid && l2 != succ {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	for l1 != mid {
		cur.Next = l1
		l1 = l1.Next
		cur = cur.Next
	}
	for l2 != succ {
		cur.Next = l2
		l2 = l2.Next
		cur = cur.Next
	}
	cur.Next = succ // 保证链表始终不断开，递归部分最终还是以succ结尾
}
