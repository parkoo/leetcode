package main

import "math"

// 思路: 二分查找, 死循环式二分查找

// 时间复杂度: O(logn)    空间复杂度: O(1)

func findPeakElement_1(nums []int) int {

	getValue := func(i int) int {
		if i == -1 || i == len(nums) {
			return math.MinInt64
		}

		return nums[i]
	}

	left, right := 0, len(nums)-1
	// 一定存在结果，死循环
	for {
		mid := left + (right-left)/2

		if getValue(mid) > getValue(mid-1) && getValue(mid) > getValue(mid+1) {
			return mid
		}

		if getValue(mid) <= getValue(mid+1) {
			left = mid + 1

		} else {
			right = mid - 1
		}
	}
}
