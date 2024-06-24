package main

import (
	"container/heap"
)

// 思路: 维护两个堆，大顶堆存储所有数据的较小的一半数据，小顶堆存储较大的一半数据
// 当数据个数为奇数时，保证大顶堆的元素个数比小顶堆多一个，此时，大顶堆的堆顶元素就是中位数
// 当数据个数为偶数时，保证大顶堆的元素个数与小顶堆的元素个数相等，此时，中位数为大顶堆的堆顶元素和小顶堆的堆顶元素的平均值
// https://leetcode.cn/problems/find-median-from-data-stream/solutions/961319/gong-shui-san-xie-jing-dian-shu-ju-jie-g-pqy8/

// 时间复杂度: addNum:O(log⁡n)，其中n为累计添加的数的数量,findMedian:O(1)   空间复杂度: O(n)

type MedianFinder struct {
	minHeap *MinHeap
	maxHeap *MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: new(MinHeap),
		maxHeap: new(MaxHeap),
	}
}

func (this *MedianFinder) AddNum(num int) {
	minHp, maxHp := this.minHeap, this.maxHeap
	if maxHp.Len() == minHp.Len() {
		if minHp.Len() == 0 || num <= minHp.Peek() {
			heap.Push(maxHp, num)

		} else {
			v := heap.Pop(minHp)
			heap.Push(maxHp, v)
			heap.Push(minHp, num)
		}

	} else {
		if num >= maxHp.Peek() {
			heap.Push(minHp, num)

		} else {
			v := heap.Pop(maxHp)
			heap.Push(minHp, v)
			heap.Push(maxHp, num)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	maxHp, minHp := this.maxHeap, this.minHeap
	if maxHp.Len() == minHp.Len() {
		return float64(maxHp.Peek()+minHp.Peek()) / 2
	}

	return float64(maxHp.Peek())
}

// 小顶堆
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() int          { return h[0] } // 返回堆顶元素

func (h *MinHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func (h *MinHeap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

// 大顶堆
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Peek() int          { return h[0] } // 返回堆顶元素

func (h *MaxHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func (h *MaxHeap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
