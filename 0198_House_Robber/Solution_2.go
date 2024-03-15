package main

// 动态规划之二  状态定义: dp[i]表示到达第i家时(可能打劫第i家也可能不打劫第i家)可以获取的全局最大金币数
// 时间复杂度：O(n)  空间复杂度：O(1)
func rob(nums []int) int {
	n := len(nums)

	// dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	pre2, pre1 := 0, 0
	res := 0
	for i := 0; i < n; i++ {
		res = max(nums[i]+pre2, pre1)
		pre2 = pre1
		pre1 = res
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
