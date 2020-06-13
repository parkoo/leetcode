package main

import (
	"math"
)

// 双指针 对撞指针，每次根据情况,移动左指针或者右指针
// 时间复杂度：O(n)  空间复杂度：O(１)

func trap_3(height []int) int {

	res := 0

	left := 0
	right := len(height)-1

	max_l := 0
	max_r := 0
	for left<right {
		if height[left]<height[right] {
			if height[left]>max_l {
				max_l = height[left]
			} else {
				res += max_l - height[left]
			}

			left++
		} else {
			if height[right]>max_r {
				max_r = height[right]
			} else {
				res += max_r - height[right]
			}

			right--
		}
	}

	return res
}