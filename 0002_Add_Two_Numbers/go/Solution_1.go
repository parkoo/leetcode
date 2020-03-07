package main

// 初等数学，两数相加
// 时间复杂度：O(max(m,n))  空间复杂度：O(max(m,n))

type ListNode struct {
    Val int
    Next *ListNode
}
 
 func addTwoNumbers_1(l1 *ListNode, l2 *ListNode) *ListNode {

	dummyHead := &ListNode{}
	cur := dummyHead

	carry := 0
	for l1!=nil && l2!=nil {
		temp := l1.Val + l2.Val + carry
		carry = temp/10
		cur.Next = &ListNode{temp%10, nil}
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1!=nil {
		temp := l1.Val + carry
		carry = temp/10
		cur.Next = &ListNode{temp%10, nil}
		cur = cur.Next
		l1 = l1.Next
	}

	for l2!=nil {
		temp := l2.Val + carry
		carry = temp/10
		cur.Next = &ListNode{temp%10, nil}
		cur = cur.Next
		l2 = l2.Next
	}

	if carry!=0 {
		cur.Next = &ListNode{carry, nil}
		cur = cur.Next
	}

	return dummyHead.Next
}