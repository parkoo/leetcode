package main

// 思路: 先连接成环, 再在相应位置断开, 最多遍历两次，相比“整体反转+部分反转”, 时间性能更好

// 时间复杂度: O(n)    空间复杂度: O(1)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight_2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 统计节点数
	cnt := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		cnt++
	}

	// 处理k
	k = k % cnt
	if k == 0 {
		return head
	}

	// 先连接成环, 再在相应位置断开
	tail := cur
	tail.Next = head
	cur = head
	for i := 0; i < cnt-k-1; i++ {
		cur = cur.Next
	}
	newHead := cur.Next
	cur.Next = nil

	return newHead
}
