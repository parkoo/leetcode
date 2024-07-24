package main

import "container/heap"

// 思路: 利用堆实现多路归并 小顶堆
// 注意去重技巧 https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/solutions/1210221/bu-chong-guan-fang-ti-jie-you-xian-dui-l-htf8/

// 时间复杂度: O(klogk)    空间复杂度: O(k)

func kSmallestPairs_1(nums1 []int, nums2 []int, k int) [][]int {
	res := make([][]int, 0)

	minHp := new(MinHeap)
	heap.Init(minHp)
	// 先将nums1[i], nums2[0] 入队，去重
	for i := 0; i < len(nums1) && i < k; i++ {
		newNode := node{
			i:   i,
			j:   0,
			sum: nums1[i] + nums2[0],
		}

		heap.Push(minHp, newNode)
	}

	cnt := 0
	for minHp.Len() > 0 && cnt < k {
		cur := heap.Pop(minHp).(node)
		curI, curJ := cur.i, cur.j
		res = append(res, []int{nums1[curI], nums2[curJ]})

		// 出堆的同时入堆
		if curJ+1 < len(nums2) {
			newNode := node{
				i:   curI,
				j:   curJ + 1,
				sum: nums1[curI] + nums2[curJ+1],
			}
			heap.Push(minHp, newNode)
		}

		cnt++
	}

	return res
}

type node struct {
	i   int // nums1[i]
	j   int // nums2[j]
	sum int // nums1[i] + nums2[j]
}

// 小顶堆
type MinHeap []node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(v interface{}) {
	*h = append(*h, v.(node))
}
func (h *MinHeap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}
