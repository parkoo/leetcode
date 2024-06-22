package main

import "container/heap"

// 思路: 哈希表+小顶堆, 维护一个大小为k的小顶堆

// 时间复杂度: O(nlogk)    空间复杂度: O(n)  cache占用空间

func topKFrequent(nums []int, k int) []int {
	// 统计出现频次
	cache := make(map[int]int)
	for _, num := range nums {
		cache[num]++
	}

	// 初始化小顶堆
	mh := new(MinHeap)
	heap.Init(mh)

	for num, freq := range cache {
		heap.Push(mh, [2]int{num, freq})
		if mh.Len() > k { // 保持堆内的元素为k个
			heap.Pop(mh)
		}
	}

	res := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(mh).([2]int)[0]
	}

	return res
}

// 实现小顶堆
type MinHeap [][2]int

func (mh MinHeap) Len() int           { return len(mh) }
func (mh MinHeap) Less(i, j int) bool { return mh[i][1] < mh[j][1] }
func (mh MinHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }

// Push and Pop use pointer receivers because they modify the slice's length, not just its contents.
func (mh *MinHeap) Push(v interface{}) {
	*mh = append(*mh, v.([2]int))
}

func (mh *MinHeap) Pop() interface{} {
	n := len(*mh)
	v := (*mh)[n-1]
	*mh = (*mh)[:n-1]
	return v
}
