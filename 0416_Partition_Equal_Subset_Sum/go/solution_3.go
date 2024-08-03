package main

// 思路: 0-1背包 物品在外层背包在内层 内层从大到小遍历

// 时间复杂度: O(n*target)    空间复杂度: O(target)

func canPartition_3(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}
	target := sum / 2

	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}

	return dp[target]
}
