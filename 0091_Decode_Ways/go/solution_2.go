package main

// 思路: 动态规划 空间降维处理

// 时间复杂度: O(n)    空间复杂度: O(1)

func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0]-'0' == 0 {
		return 0
	}

	// dp[i] = k, 表示以i结尾的字符串可以表示的解码总数
	// dp[i] = dp[i-1] + dp[i-2] (s[i] 和 s[i-1]组合能够正确解码)
	cur := 1           // 初始化 dp[0] = 1
	pre1, pre2 := 0, 1 // pre1代表dp[i-2], pre2代表dp[i-1]

	for i := 1; i < n; i++ {
		cur = 0 // dp[i] 初始值为0

		// s[i]能够正确解码
		if s[i]-'0' != 0 {
			cur = pre2
		}

		// s[i-1]和s[i]的组合能够正确解码
		if (s[i-1]-'0')*10+(s[i]-'0') >= 10 && (s[i-1]-'0')*10+(s[i]-'0') <= 26 {
			if i > 1 {
				cur += pre1

			} else { // 开头的前两位s[0]s[1]
				cur += 1
			}
		}

		pre1 = pre2
		pre2 = cur
	}

	return cur
}
