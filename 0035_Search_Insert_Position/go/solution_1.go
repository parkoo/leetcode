package main

// 思路: 二分查找

// 时间复杂度: O(log(n))    空间复杂度: O(1)

func searchInsert(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if target > nums[n-1] {
		return n
	}

	l, r := 0, n-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
