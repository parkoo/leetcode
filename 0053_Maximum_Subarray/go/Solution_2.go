package main

// 思路：降维动态规划  注意和 “lc152求积最大的子数组” 比较，求积会受符号变化的影响，需要同时记录最大值和最小值

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
