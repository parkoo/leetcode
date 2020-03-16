
package main

// 递归法,假设当前节点的子节点下的节点已经完成反转，需反转当前节点与当前节点的子节点
// 时间复杂度：O(n)  空间复杂度：O(n)

 /**
  type ListNode struct {
      Val int
      Next *ListNode
   }
*/

 func reverseList_2(head *ListNode) *ListNode {
	cur := head

	if cur==nil || cur.Next==nil {
		return cur
	}

	// 1. 完成当前节点cur的子节点下的节点的反转
	// 即完成　pre1 -> pre2 -> cur -> next1 <- next2
	p := reverseList_2(cur.Next)  

	// 2. 反转当前节点和当前节点的子节点
	                       //                     ->
	cur.Next.Next = cur    // pre1 -> pre2 -> cur    next1 <- next2
	                       //                     <-

	cur.Next = nil        // pre1 -> pre2 -> cur <- next1 <- next2
	
	return p
}