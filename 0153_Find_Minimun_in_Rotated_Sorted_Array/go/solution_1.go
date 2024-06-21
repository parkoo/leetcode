package main

// 思路: 二分查找
// 我们考虑数组中的最后一个元素"x": 在最小值右侧的元素（不包括最后一个元素本身），它们的值一定都严格小于"x",
// 而在最小值左侧的元素，它们的值一定都严格大于"x"。因此，我们可以根据这一条性质，通过二分查找的方法找出最小值

// 时间复杂度: O(log(n))    空间复杂度: O(1)

func findMin_1(nums []int) int {
	n := len(nums)

	flag := nums[n-1] // 以最后一个元素作为标记
	l, r := 0, n-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] <= flag { // nums[mid] < flag 也可
			r = mid
		} else {
			l = mid + 1
		}
	}

	return nums[l]
}
