package main

// Solution_1的简化版，初等数学，两数相加
// 时间复杂度：O(max(m,n))  空间复杂度：O(max(m,n))

// type ListNode struct {
//     Val int
//     Next *ListNode
// }
 
 func addTwoNumbers_2(l1 *ListNode, l2 *ListNode) *ListNode {

	dummyHead := &ListNode{}
	cur := dummyHead

	carry := 0
	for l1!=nil || l2!=nil {
		x := 0
		y := 0
		if l1!=nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2!=nil {
			y = l2.Val
			l2 = l2.Next		
		}
		temp := x + y + carry
		carry = temp/10
		cur.Next = &ListNode{temp%10, nil}
		cur = cur.Next
	}

	if carry != 0 {
		cur.Next = &ListNode{carry, nil}
		cur = cur.Next
	}

	return dummyHead.Next
}