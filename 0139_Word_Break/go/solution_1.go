package main

// 思路: 动态规划
// 定义 dp[s] 表示字符串 s 是否可以被拆分,
// 因为 s = (s - word) + word, 所以，若 dp[s-word] == true (range work in wordDict), 则 dp[s] = true

// 时间复杂度: O(s*n)    空间复杂度: O(n)

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make(map[string]bool)

	dp[""] = true // 前置初始状态
	for i := 1; i <= n; i++ {
		cur := s[:i]
		for _, word := range wordDict {
			if i >= len(word) && s[i-len(word):i] == word {
				pre := strSub(cur, word)
				if dp[pre] {
					dp[cur] = true
					break
				}
			}
		}
	}

	return dp[s]
}

func strSub(a, b string) string {
	if len(a) <= len(b) {
		return ""
	}

	return a[:len(a)-len(b)]
}
