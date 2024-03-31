package main

// 思路: 动态规划 二维空间
// dp[i][j] 表示考虑（可选放入也可选择不放入）将第i个物品放入容量为j的背包，是否可以恰好填满容量为j的背包

// 时间复杂度: O(n*target)    空间复杂度: O(n*target)

func canPartition_1(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make([][]bool, len(nums))
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}

	for i := 0; i <= target; i++ {
		dp[0][i] = nums[0] == i
	}
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		for j := 0; j <= target; j++ {
			dp[i][j] = dp[i-1][j] // j < num 时，背包容量不够放入当前物品，只能选择不放入当前物品，即需要先初始化为取上一个物品的状态
			if j >= num {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-num]
			}
		}
	}

	return dp[len(nums)-1][target]
}
