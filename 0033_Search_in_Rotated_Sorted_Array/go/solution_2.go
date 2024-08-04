package main

// 思路: 二分查找

// 时间复杂度：O(lgn)  空间复杂度：O(1)

func search_2(nums []int, target int) int {
	l, r := 0, len(nums)-1
	flag := nums[0]

	for l < r {
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
