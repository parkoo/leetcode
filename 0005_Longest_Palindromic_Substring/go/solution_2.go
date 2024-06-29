package main

// 思路：动态规划  dp[i][j] == true 表示 s[i,j]为回文串

// 时间复杂度：O(n^2)  空间复杂度：O(n^2)

func longestPalindrome_2(s string) string {
	n := len(s)

	// dp[i][j] == true 表示 s[i,j]为回文串
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	maxLen := 1   // 记录最长回文串的长度
	res := s[0:1] // 回文串

	// dp[i][j] 依赖dp[i+1][j-1] 外层为end, 内层为start
	for end := 1; end < n; end++ {
		for start := 0; start < end; start++ {
			// end - start == 0, 表示此时只针对一个字符，必定是回文串 如：a
			// end - start == 1, 表示此时针对两个字符，若两个字符相等则是回文串 如：aa
			// end - start == 2, 表示此时针对三个字符，中间的字符不用管，两边的字符相等即为回文串 如：aba
			// 其余超过三个字符的（end - start > 2）的情况，需根据前一个状态dp[start+1][end-1]来判断
			if (end-start <= 2 || dp[start+1][end-1]) && s[start] == s[end] {
				dp[start][end] = true
				if end-start+1 > maxLen {
					maxLen = end - start + 1
					res = s[start : end+1]
				}
			}
		}
	}

	return res
}
