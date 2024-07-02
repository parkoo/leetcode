package main

// 思路：双指针  快慢指针

// 时间复杂度：O(n)  空间复杂度：O(1)

func removeDuplicates_1(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	// 保证nums[0,i]区间内为重复个数不超过2个的数字
	i, j := 1, 2
	for ; j < n; j++ {
		if nums[j] > nums[i-1] {
			i++
			nums[i] = nums[j]
		}

	}

	return i + 1
}
