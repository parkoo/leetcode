package main

import (
	"sort"
)

// 贪心思路
// 时间复杂度：O(nlogn)  空间复杂度：O(1)

func findBestValue_2(arr []int, target int) int {
	sort.Ints(arr)

	n := len(arr)
	sum := 0
	for i:=0; i<n; i++ {
		cur := (target-sum)/(n-i)
		if cur <= arr[i] {
			doubleCur := (float64(target)-float64(sum))/float64(n-i)
			if doubleCur-float64(cur)<=0.5 {
				return cur
			}else {
				return cur+1
			}
		}

		sum += arr[i]
	}

	return arr[n-1]
}