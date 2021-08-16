package main

import "sort"

// 排序后，将数组平均分成两部分，则后一部分任意一个元素均大于前一部分任一元素，则后一部分任意一个元素均不是前一部分任意两个元素和的一半
// 时间复杂度：O(nlgn)  空间复杂度：O(1)

func rearrangeArray_1(nums []int) []int {
	sort.Slice(nums, func(i int, j int) bool {
		return nums[i] < nums[j]
	})

	m := (len(nums) - 1) / 2
	j := m + 1
	if m%2 == 0 {
		j++
	}
	for i := 1; i <= m; i += 2 {
		nums[i], nums[j] = nums[j], nums[i]
		j += 2
	}

	return nums
}
