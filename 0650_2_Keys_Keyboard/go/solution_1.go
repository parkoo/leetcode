package main

// 思路: 动态规划

// 时间复杂度: O(n^2)    空间复杂度: O(n)

func minSteps_1(n int) int {

	// 定义 dp[i] = k, 表示生成i个字符需要的最少操作次数为k
	// 对于长度为i的字符串，若i是j的整数倍, 即（i%j==0, j<i）,
	// 则长度为i的字符串可以由长度为j的字符经过1次拷贝以及（i/j-1）次粘贴操作等到,共（i/j）次操作， 即 dp[i] = dp[j] + (i/j)
	// 从中选取操作次数最小的一种即可得到dp[i]

	// 例如：
	// dp[12] =  dp[1] + 12 (对"A"进行1次拷贝后再进行11次粘贴)
	// dp[12] =  dp[2] + 6  (对"AA"进行1次拷贝后再进行5次粘贴)
	// dp[12] =  dp[3] + 4  (对"AAA"进行1次拷贝后再进行3次粘贴)
	// dp[12] =  dp[4] + 3  (对"AAAA"进行1次拷贝后再进行2次粘贴)
	// dp[12] =  dp[6] + 2  (对"AAAAAA"进行1次拷贝后再进行1次粘贴)
	dp := make([]int, n+1)

	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = i // j = 1时， dp[i] = i

		// 这里可以做一些剪枝，当 j > i/2 时， i % j != 0
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				dp[i] = min(dp[i], dp[j]+(i/j))
			}
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
