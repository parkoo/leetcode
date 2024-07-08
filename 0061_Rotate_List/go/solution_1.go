package main

// 思路: 先整体反转，再部分反转

// 时间复杂度: O(n)    空间复杂度: O(1)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight_1(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 处理k
	cnt := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		cnt++
	}
	k = k % cnt
	if k == 0 || k == cnt {
		return head
	}

	// 整体反转
	start, end := head, head
	for end.Next != nil {
		end = end.Next
	}
	curHead, _ := reverse(start, end)
	dummyHead := new(ListNode)
	dummyHead.Next = curHead

	// 反转前半部分
	start, end = dummyHead.Next, dummyHead
	for i := 0; i < k; i++ {
		end = end.Next
	}
	next := end.Next
	curHead, curTail := reverse(start, end)
	dummyHead.Next = curHead
	curTail.Next = next

	// 反转后半部分
	pre := curTail
	start, end = next, next
	for end.Next != nil {
		end = end.Next
	}
	curHead, _ = reverse(start, end)
	pre.Next = curHead

	return dummyHead.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	var dummyHead *ListNode
	tail.Next = nil
	pre := dummyHead
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return tail, head
}
