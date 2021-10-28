package main

// 虚拟头节点的使用
// 时间复杂度：O(n)  空间复杂度：O(1)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	// 虚拟头节点
	dummyHead := new(ListNode)
	dummyHead.Next = head
	cur := dummyHead

	for cur.Next != nil {
		if cur.Next.Val == val {
			node := cur.Next
			cur.Next = cur.Next.Next
			node.Next = nil
		} else {
			cur = cur.Next
		}
	}

	return dummyHead.Next
}
