package main

// 思路: 划分翻转区间，逐个翻转区间

// 时间复杂度: O(n) （其中 head 指针会在n/k个节点上停留，每次停留需要进行一次k次翻转）  空间复杂度: O(1)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup_1(head *ListNode, k int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head

	// pre为待翻转区间的前驱节点，cur为当前待翻转区间的起始节点
	pre, cur := dummyHead, head
	for cur != nil {
		// 通过pre确定当前待翻转区间的尾节点tail
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil { // 待翻转的节点个数不足k个，直接返回
				return dummyHead.Next
			}
		}

		// 记录下一个待翻转区间的起始位置
		next := tail.Next

		// 翻转当前区间节点
		curHead, curTail := reverse(cur, tail)
		// 连接前驱节点和后续节点
		pre.Next = curHead
		curTail.Next = next

		// 迭代到下一个翻转区间
		pre = curTail
		cur = next
	}

	return dummyHead.Next
}

// 翻转 [head, tail] 之间的节点，返回新的head、tail
func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	dummyHead := new(ListNode)
	dummyHead.Next = head
	tail.Next = nil

	pre := dummyHead
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	// 此时的pre为翻转后的head, 此时的head为翻转后的tail
	return pre, head
}
