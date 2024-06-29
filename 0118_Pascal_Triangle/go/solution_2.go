package main

// 思路：一维动态规划, 除去结果所占空间需要O(n)的空间

// 时间复杂度：O(n^2)  空间复杂度：O(n) 不包含结果存储

func generate_2(numRows int) [][]int {
	res := make([][]int, 0)

	dp := make([]int, numRows)
	dp[0] = 1
	res = append(res, []int{1})
	for i := 1; i < numRows; i++ {
		// dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
		// 当前位置依赖上一行同位置以及上一行前一个位置，防止上一行前一个位置被覆盖需要从后向前遍历
		for j := numRows - 1; j >= 0; j-- {
			dp[j] = dp[j]
			if j >= 1 {
				dp[j] = dp[j] + dp[j-1]
			}
		}

		item := make([]int, 0)
		for _, v := range dp {
			if v != 0 {
				item = append(item, v)
			}
		}
		res = append(res, item)
	}

	return res
}
