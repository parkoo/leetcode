package main

// 思路: 0-1背包最值问题
// 物品在外层且内层从大到小遍历

// 注意和lc0322的区别，其为完全背包最值问题，物品可以在外层也可以在内层，物品在外层时，内层从小到大遍历，保证物品可以重复使用
// 而本题为0-1背包最值问题，需物品在外层且内层从大到小遍历，以保证每个物品只使用一次

// 时间复杂度: O(m*n*len(str))    空间复杂度: O(m*n)

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1

	for _, str := range strs {
		zeroCnt, oneCnt := cntZeroAndOne(str)
		for i := m; i >= zeroCnt; i-- {
			for j := n; j >= oneCnt; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeroCnt][j-oneCnt]+1)
			}
		}
	}

	return dp[m][n]
}

func cntZeroAndOne(s string) (int, int) {
	n := len(s)
	cntZero := 0
	for i := 0; i < n; i++ {
		if int(s[i]-'0') == 0 {
			cntZero++
		}
	}

	return cntZero, n - cntZero
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
