package main

// 思路: 优化动态规划
// 优化: 只考虑 j==2 和 j==3的情况即可
// 推导：如果拆分的因子中有一个为a(a >= 4), 则可以把a进一步拆分成2、a-2, 因为2*(a-2) = 2*a-4 >= a (a>=4)
// 即 为了达到积最大的效果被拆分的因子必定只能是2或3， 因为对于大于等于4的因子a, 进一步拆分成2*(a-2)会使得效果更优(积更大)

// 时间复杂度: O(n)    空间复杂度: O(n)

func integerBreak(n int) int {

	// dp[i] = k, 表示将i拆分成至少2个正整数后,这些正整数的最大积为k
	// 将i拆分成两部分j和(i-j)
	// 为了保证至少拆分成2个正整数，可以分为两种情况
	// 1. j不拆分且(i-j)也不拆分，此时 dp[i] = j*(i-j)
	// 2. j不拆分但(i-j)作拆分， 此时 dp[i] = j*dp[i-j] (dp[i-j]表示把(i-j)至少拆分成两部分的最大积)
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		// 优化: 只考虑 j==2 和 j==3的情况即可
		for j := 2; j <= 3; j++ {
			dp[i] = max(max(dp[i], j*dp[i-j]), j*(i-j))
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
