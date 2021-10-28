package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 要删除链表中的一个节点，需要拿到它的前驱节点;
// 此题中无法取得待删除节点的前一个节点，思路是将待删除节点的后一个节点的值赋给待删除节点，然后删除待删除节点的后一个节点.
// 时间复杂度：O(1)  空间复杂度：O(1)

func deleteNode(node *ListNode) {
	next := node.Next
	node.Val = next.Val
	node.Next = node.Next.Next
	next.Next = nil
}
