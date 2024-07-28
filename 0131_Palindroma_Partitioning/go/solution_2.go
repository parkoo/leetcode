package main

// 思路：回溯法 + 动态规划预处理

func partition(s string) [][]string {
	res := make([][]string, 0)

	// 动态规划 预处理回文串 方便搜索
	// dp[i][j] == true, 表示 s[i:j+1] 为回文串
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}
	dp[0][0] = true

	// dp[i][j] 依赖dp[i+1][j-1] 外层为end, 内层为start
	for end := 1; end < len(s); end++ {
		// 从大到小遍历
		for start := end; start >= 0; start-- {
			if (end-start <= 2 || dp[start+1][end-1]) && s[start] == s[end] {
				dp[start][end] = true
			}
		}
	}

	// 回溯切分字符串
	var backtrack func(s string, start int, sub []string)
	backtrack = func(s string, start int, sub []string) {
		if start == len(s) {
			item := make([]string, 0)
			item = append(item, sub...)
			res = append(res, item)
			return
		}

		for i := 1; i <= len(s)-start; i++ {
			end := start + i
			cur := s[start:end]

			if dp[start][end-1] { // 判断是否为回文串
				sub = append(sub, cur)
				backtrack(s, end, sub)
				sub = sub[:len(sub)-1]
			}
		}
	}

	backtrack(s, 0, []string{})
	return res
}
