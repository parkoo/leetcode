package main

import "sort"

// 思路: 链表转化为数组，排序，再转化为链表

// 时间复杂度: O(nlogn)    空间复杂度: O(n)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList_1(head *ListNode) *ListNode {
	arr := make([]int, 0)
	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	cur = head
	i := 0
	for cur != nil {
		cur.Val = arr[i]
		i++
		cur = cur.Next
	}

	return head
}
