package main

// 思路： 双指针 快慢指针 同solution_1

// 时间复杂度：O(n)  空间复杂度：O(1)

func removeDuplicates_2(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	left := 0
	for right := 1; right < n; right++ {
		if nums[right] != nums[left] {
			left++
			nums[right], nums[left] = nums[left], nums[right]
		}
	}

	return left + 1
}
