package main

import "container/heap"

// 堆的应用, 可以以lgn的复杂度获取数组中的最值
// 时间复杂度：O(klgn) 空间复杂度：O(1)

func minStoneSum_1(piles []int, k int) int {
	in := IntArr(piles)
	heap.Init(&in)
	for i := 0; i < k; i++ {
		v := heap.Pop(&in)
		vv, _ := v.(int)
		vv = vv - vv/2
		heap.Push(&in, vv)
	}

	return in.Sum()
}

type IntArr []int

func (arr IntArr) Len() int           { return len(arr) }
func (arr IntArr) Less(i, j int) bool { return arr[i] > arr[j] } // 大顶堆
func (arr IntArr) Swap(i, j int)      { arr[i], arr[j] = arr[j], arr[i] }

func (arr *IntArr) Push(v interface{}) {
	*arr = append(*arr, v.(int))
}

func (arr *IntArr) Pop() interface{} {
	old := *arr
	n := len(old)
	v := old[n-1]
	*arr = old[:n-1]
	return v
}

func (arr IntArr) Sum() int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}
