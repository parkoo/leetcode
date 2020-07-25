package main

import (
	"math"
)

// 动态规划
// dp[i][j]表示把 i 个数划分成 j 个区间时的最大连续区间和的最小值.
// 在进行状态转移时，考虑第 j 段的范围，即可以枚举 k，其中前 k 个数被分割为前 j−1 段，而第 k+1 到第 i 个数为第 j 段．
// dp[i][j] = Min(dp[i][j], Max(dp[k][j-1], Sum(nums[k+1],nums[k+2],...,nums[i]))) for (0<=k<i)
// 对于dp[i][j]，由于不能分出空的子数组，因此合法的状态必须有 i≥j。对于不合法（i<j）的状态，
// 由于我们的目标是求出最小值，因此可以将这些状态全部初始化为一个很大的数。
// 此外，还需要将 dp[0][0] 的值初始化为 0。在上述的状态转移方程中，当 j=1 时，唯一的可能性就是前 i 个数被分成了一段.
// 如果枚举的 k=0，那么就代表着这种情况；如果 k!=0，对应的状态 dp[k][0] 是一个不合法的状态，无法进行转移。
// 因此我们需要令 dp[0][0] = 0。

// 时间复杂度：O(n^2*m)  空间复杂度：O(n*m)

func splitArray_1(nums []int, m int) int {

	// 构造动态规划区间，并初值化
	n := len(nums)
	dp := make([][]int, n+1)
	for i:=0; i<len(dp); i++ {
		dp[i] = make([]int, m+1)
		for j:=0; j<len(dp[i]); j++ {
			dp[i][j] = math.MaxInt32    // 数值初始化
		}
	}
	dp[0][0] = 0  // 数值初始化

	// 构造辅助区间和，sum[i]表示前i个数的和
	sum := make([]int, n+1)
	for i:=0; i<n; i++ {
		sum[i+1] = sum[i] + nums[i]
	}

	// 状态转移
	for i:=1; i<=n; i++ {
		for j:=1; j<=int(math.Min(float64(m),float64(i))); j++ {
			for k:=0; k<i; k++ {
				dp[i][j] = int(math.Min(float64(dp[i][j]), math.Max(float64(dp[k][j-1]),float64(sum[i]-sum[k]))))
			}
		}
	}

	return dp[n][m]
}