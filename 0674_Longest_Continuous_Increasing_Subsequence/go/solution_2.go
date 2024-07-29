package main

// 思路： 动态规划

// 对 solution_1 进行空间上的降维

// 时间复杂度：O(n)  空间复杂度：O(1)

func findLengthOfLCIS_2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	res := 1
	pre, cur := 1, 1
	for i := 1; i < n; i++ {
		cur = 1
		if nums[i] > nums[i-1] {
			cur = pre + 1
		}
		pre = cur

		if cur > res {
			res = cur
		}
	}

	return res
}
