package main

// 动态规划
// 对于正整数x, 将其二进制表示bit[x] 右移以为则可得到x/2的二进制表示bit[x/2], 故，得知bit[x/2], 便可求得bit[x]
// 对于奇数x, bit[x] = bit[x/2] + 1
// 对于偶数x, bit[x] = bit[x/2]
// 时间复杂度：O(n)  空间复杂度：O(1)

func countBits_1(n int) []int {
	// dp[i]表示整数 i 二进制表示中 1 的个数
	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		dp[i] = dp[i/2]
		if i%2 == 1 {
			dp[i] += 1
		}
	}

	return dp
}
