package main

// 双指针  快慢指针
// 时间复杂度：O(n)  空间复杂度：O(1)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	// 保证nums[0,i]区间内为重复个数不超过2个的数字
	i, j := 1, 2
	for ; j < len(nums); j++ {
		if nums[j] > nums[i-1] {
			i++
			nums[i] = nums[j]
		}

	}

	return i + 1
}
