package main

// 思路: 二分查找  以nums[l]为标定点

// 时间复杂度：O(lgn)  空间复杂度：O(1)

func search_3(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l < r {
		// 注意这里以nums[l]为标定点 而不是nums[0]
		flag := nums[l]

		m := l + (r-l)/2
		if nums[m] >= flag {
			if target >= flag && target <= nums[m] {
				r = m
			} else {
				l = m + 1
			}

		} else {
			if target < flag && target > nums[m] { // 这里是大于而非大于等于
				l = m + 1
			} else {
				r = m
			}
		}
	}

	if nums[l] == target {
		return l
	}

	return -1
}
