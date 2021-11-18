package main

// 双指针
// 时间复杂度：O(n)  空间复杂度：O(1)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next

		// 找到快慢指针相交点
		if s == f {
			// 从起点和相交点同时同速出发，最终在环的入口相交
			k := head
			for s != k {
				s = s.Next
				k = k.Next
			}
			return k
		}
	}

	return nil
}
