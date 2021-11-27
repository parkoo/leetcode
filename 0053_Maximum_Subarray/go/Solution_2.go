package main

// 降维动态规划
// 时间复杂度：O(n)  空间复杂度：O(1)

func maxSubArray_2(nums []int) int {
	n := len(nums)

	dp := nums[0]
	res := nums[0]
	for i := 1; i < n; i++ {
		dp = max(nums[i], nums[i]+dp)
		res = max(res, dp)
	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
