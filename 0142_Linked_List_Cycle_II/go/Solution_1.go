package main

// 双指针
// 时间复杂度：O(n)  空间复杂度：O(1)

// 令快指针的速度是慢指针的2倍, 设一定时间后，慢指针走过的路径为l,快指针走过的路径为2l
// 若路径中有环，则快慢指针一定会在环内相遇(此时快指针甩开慢指针n圈)
// 设圈外的长度为d,一圈的长度为c,快慢指正在环内相遇时,慢指针在环内走的长度为a,剩余的长度为b，则快指针在环内走了n圈加a
// l = d + a
// 2l = d + n *（a + b）+ a = d + (n + 1) * a + n * b
// 所以，2 * (d + a) = d + (n + 1) * a + n * b， 即，d = (n - 1) * (a + b) + b
// 即 圈外的长度等于从相遇点走(n-1)圈后，再走到入圈点

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
