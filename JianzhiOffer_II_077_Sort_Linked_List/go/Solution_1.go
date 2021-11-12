package main

import (
	"container/heap"
)

// 使用堆
// 时间复杂度：O(nlogn)  空间复杂度：O(n)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {

	items := Items([]ListNode{})
	heap.Init(&items)

	cur := head
	for cur != nil {
		v := *cur
		heap.Push(&items, v)
		cur = cur.Next
	}

	dummyHead := &ListNode{}
	cur = dummyHead
	for items.Len() > 0 {
		v := heap.Pop(&items).(ListNode)
		v.Next = nil // 避免原Next对结果的影响
		cur.Next = &v
		cur = cur.Next
	}

	return dummyHead.Next
}

type Items []ListNode

func (items Items) Len() int               { return len(items) }
func (items Items) Less(i int, j int) bool { return items[i].Val < items[j].Val }
func (items Items) Swap(i int, j int)      { items[i], items[j] = items[j], items[i] }
func (items *Items) Push(v interface{}) {
	*items = append(*items, v.(ListNode))
}
func (items *Items) Pop() interface{} {
	old := *items
	v := old[len(old)-1]
	*items = old[:len(old)-1]
	return v
}
