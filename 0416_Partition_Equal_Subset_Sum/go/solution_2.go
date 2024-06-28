package main

// 思路: 动态规划 一维空间
// dp[i][j] 表示考虑（可选放入也可选择不放入）将第i个物品放入容量为j的背包，是否可以恰好填满容量为j的背包
// 进一步优化为一维空间

// 时间复杂度: O(n*target)    空间复杂度: O(target)

func canPartition_2(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make([]bool, target+1)
	for i := 0; i <= target; i++ {
		if i == nums[0] {
			dp[i] = true
		}
	}

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		// dp[i][j] = dp[i-1][j] || dp[i-1][j-num]
		// 对于二维dp来说，dp[i][j] 依赖 dp[i-1][j]、 dp[i-1][j-num], 从右侧遍历可以避免dp[i-1][j-num]被覆盖
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}
