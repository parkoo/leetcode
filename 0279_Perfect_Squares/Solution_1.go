package main

import "math"

// 思路：完全背包最值问题，从一系列物品中(完全平方数)选择一些填大小为n的背包
// 解法1： 背包在外层物品在内层
// 动态规划：dp[i]表示最少需要多少个完全平方数来表示数字i,
// 因为 i = j^2 + (i - j^2), 所以，构成 "i" 所需要的平方数的个数, 可以看做是通过构成 "i-j^2" 的平方数的个数再加上一个平方数（j^2）求得
// 即：p[i] = 1 + min(dp[i-j^2]) when 1<=j^2<=i

// 时间复杂度：O(n)  空间复杂度：O(n)

func numSquares_1(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
		for j := 1; j*j <= i; j++ {
			dp[i] = min(1+dp[i-j*j], dp[i])
		}
	}

	return dp[n]
}
