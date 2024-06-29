package main

// 思路: 快慢指针将链表均分成两部分，翻转后半部分链表，双指针判断是否是回文

// 时间复杂度: O(n)    空间复杂度: O(1)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome_2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 快慢指针将链表均分成两部分，若为奇数个节点，前半部分多一个
	// fast取head.Next 而不是head
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	newHead := slow.Next
	slow.Next = nil

	// 翻转后半部分链表
	var pre *ListNode // 这里只可申明变量，不可new一个出来，new一个出来时有val值的
	cur := newHead
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	// 双指针判断是否是回文
	cur1, cur2 := head, pre
	for cur1 != nil && cur2 != nil {
		if cur1.Val != cur2.Val {
			return false
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	return true
}
