package main

// 思路: 链表 归并排序 自顶向下递归式

// 时间复杂度: O(nlgn)    空间复杂度: O(lgn) 递归栈空间

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList_2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 快慢指针 将链表均分成两部分
	// 这里fast必须取head.Next, 若取了head, 则在链表节点为偶数时，切分不均匀，导致时间复杂度和空间复杂度上升
	// 或者这里先加个虚拟头结点dummyHead, 可以令slow, fast := dummyHead, dummyHead, 效果也是均分的
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	left, right := head, slow.Next
	slow.Next = nil // 断成两部分

	leftList := sortList_2(left)
	rightList := sortList_2(right)
	return merge(leftList, rightList)
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dummyHead := new(ListNode)
	cur := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next

		} else {
			cur.Next = l2
			l2 = l2.Next
		}

		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return dummyHead.Next
}
