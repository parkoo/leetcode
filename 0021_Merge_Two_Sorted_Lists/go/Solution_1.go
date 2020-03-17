package main

// 迭代法,归并过程
// 时间复杂度：O(m+n)  空间复杂度：O(1)  其中，m和n分别是两个链表的长度

 type ListNode struct {
     Val int
     Next *ListNode
 }

 func mergeTwoLists_1(l1 *ListNode, l2 *ListNode) *ListNode {

	 dummy_head := &ListNode{}
	 cur := dummy_head

	 for l1!=nil && l2!=nil {
		 if l1.Val<l2.Val {
			 cur.Next = l1
			 l1 = l1.Next
		 }else {
			 cur.Next = l2	 
			 l2 = l2.Next
		 }
		 cur = cur.Next
	 }

	 if l1==nil {
		 cur.Next = l2
	 }
	 if l2==nil {
		 cur.Next = l1
	 }

	 return dummy_head.Next
}