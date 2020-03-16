package main

// 迭代法
// 时间复杂度：O(n)  空间复杂度：O(1)

 type ListNode struct {
    Val int
    Next *ListNode
 }

 func reverseList_1(head *ListNode) *ListNode {

	 var pre *ListNode
	 cur := head

	 for cur!=nil {
		 next := cur.Next
		 cur.Next = pre
		 pre = cur
		 cur = next
	 }

	 return pre
}