package main

import (
	"sort"
	"math"
)

// 二分查找
// 时间复杂度：O(nlogn)  空间复杂度：O(n)

func findBestValue_1(arr []int, target int) int {
	sort.Ints(arr)

	// 前缀和sum[i] 表示 arr[0]+arr[1]+...+arr[i], 可简化后续运算
	n := len(arr)
	sum := make([]int, n)
	sum[0] = arr[0]
	for i:=1; i<n; i++ {
		sum[i] = sum[i-1] + arr[i]
	}
 
	// 二分查找
	l := 0
	r := arr[n-1]
	for l<r {
		mid := l+(r-l)/2
 
		index := sort.SearchInts(arr, mid)	 
		var cur int
		if index == 0 {
			cur = n * mid
		}else {
			cur = sum[index-1] + (n-index)*mid
		}
		 
		if cur>=target {
			r = mid
		}else {
			l = mid + 1
		} 
	}
 
	// 在 l 和（l-1）中确定出结果最接近target的最小的值
	curIndex := sort.SearchInts(arr, l)
	var curSum int
	if curIndex == 0 {
		curSum = n * l
	}else {
		curSum = sum[curIndex-1] + (n-curIndex)*l
	}
 
	preIndex := sort.SearchInts(arr, l-1)
	var preSum int
	if preIndex == 0 {
		preSum = n * (l-1)
	}else {
		preSum = sum[preIndex-1] + (n-preIndex)*(l-1)
	}
 
	if int(math.Abs(float64(preSum-target))) <= int(math.Abs(float64(curSum-target))){
		return l-1
	}else {
		return l
	}
}
