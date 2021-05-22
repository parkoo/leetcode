package main

import (
	"container/heap"
)

// 哈希表去重 + 优先队列取前K值
// 时间复杂度：O(l*n + l*mlgk), 空间复杂度：O(l*m + l*K) ,l为字符串平均长度，n为给定字符串数组长度，m为去重后字符串数组长度

func topKFrequent_1(words []string, k int) []string {
	wfMap := make(map[string]int, 0)
	for _, word := range words {
		wfMap[word]++
	}

	hp := &hp{}
	for w, f := range wfMap {
		heap.Push(hp, pair{w, f})
		if hp.Len() > k {
			heap.Pop(hp)
		}
	}

	res := make([]string, k)
	for i := len(res) - 1; i >= 0; i-- {
		res[i] = heap.Pop(hp).(pair).word
	}
	return res
}

type pair struct {
	word string
	freq int
}

// 最小堆实现
type hp []pair

func (h hp) Len() int {
	return len(h)
}
func (h hp) Less(i, j int) bool {
	return h[i].freq < h[j].freq || h[i].freq == h[j].freq && h[i].word > h[j].word
}
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(pair))
}
func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
