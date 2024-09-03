package main

import (
	"math"
	"strconv"
)

// 思路:  下一个排列  类似lc0031

// 时间复杂度: O(lgn) (N = log(10)n，N是n转成字符串后的长度序列。而log以10为底和以2为底在渐进意义上没有区别)   空间复杂度: O(lgn)

func nextGreaterElement(n int) int {
	nums := []byte(strconv.Itoa(n))

	i, j := len(nums)-2, len(nums)-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i < 0 {
		return -1
	}

	k := len(nums) - 1
	if i >= 0 {
		for nums[k] <= nums[i] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}

	l, r := j, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

	res, _ := strconv.Atoi(string(nums))
	if res > math.MaxInt32 {
		return -1
	}
	return res
}
