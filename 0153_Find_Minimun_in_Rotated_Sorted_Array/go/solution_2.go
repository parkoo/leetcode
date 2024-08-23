package main

// 思路: 二分查找 以nums[r]作为标定点

// 时间复杂度: O(log(n))    空间复杂度: O(1)

func findMin_1(nums []int) int {
	n := len(nums)

	l, r := 0, n-1
	for l < r {
		flag := nums[r] // 以nums[r]作为标记
		mid := l + (r-l)/2
		if nums[mid] <= flag { // nums[mid] < flag 也可
			r = mid
		} else {
			l = mid + 1
		}
	}

	return nums[l]
}
