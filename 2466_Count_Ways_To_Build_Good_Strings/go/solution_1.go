package main

// 思路: 动态规划
// 类似lc0091
// 类似爬楼梯问题  dp[i]可以由dp[i-1]演变而来，也可以由dp[i-2]演变而来,通过条件由dp[i-1]和dp[i-2]来获得dp[i]

// 时间复杂度: O(high)    空间复杂度: O(high)

func countGoodStrings(low int, high int, zero int, one int) int {
	// dp[i]=k 表示构造长为 i 的字符串的方案数为k
	// dp[i] 可以由dp[i-zero]构建而来，也可以由dp[i-one]构建而来
	// dp[i] = dp[i-zero] + dp[i-one]
	dp := make([]int, high+1)
	dp[0] = 1 // 构造空串的方案数为 1

	for i := 1; i <= high; i++ {
		dp[i] = 0
		if i >= zero {
			dp[i] = dp[i-zero] % (1000000007)
		}
		if i >= one {
			dp[i] = (dp[i] + dp[i-one]) % (1000000007)
		}
	}

	res := 0
	for i := low; i <= high; i++ {
		res = (res + dp[i]) % (1000000007)
	}

	return res
}
