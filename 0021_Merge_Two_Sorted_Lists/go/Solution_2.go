package main

// 递归法
// 时间复杂度：O(m+n)  空间复杂度：O(m+n)  其中，m和n分别是两个链表的长度

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
 func mergeTwoLists_2(l1 *ListNode, l2 *ListNode) *ListNode {

	 if l1==nil {
		 return l2
	 }
	 if l2==nil {
		 return l1
	 }

	 if l1.Val<l2.Val {
		 l1.Next = mergeTwoLists_2(l1.Next, l2)
		 return l1
	 }else {
		 l2.Next = mergeTwoLists_2(l1, l2.Next)
		 return l2
	 }
}