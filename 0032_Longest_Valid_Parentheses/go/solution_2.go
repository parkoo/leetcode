package main

// 思路: 动态规划
// https://leetcode.cn/problems/longest-valid-parentheses/solutions/206995/dong-tai-gui-hua-si-lu-xiang-jie-c-by-zhanganan042/?envType=study-plan-v2&envId=top-100-liked

// 时间复杂度: O(n)    空间复杂度: O(n)

func longestValidParentheses(s string) int {
	res := 0
	if len(s) == 0 {
		return res
	}

	// dp[i] 表示以 s[i] 为结尾的最长有效括号的长度
	dp := make([]int, len(s))
	dp[0] = 0
	for i := 1; i < len(s); i++ {
		// s[i] 为 '('
		if s[i] == '(' {
			dp[i] = 0 // 默认为0，该行可省略
			continue
		}

		// s[i] 为 ')' 时，分两种情况
		// 1. s[i-1] 为 '('
		if s[i-1] == '(' {
			dp[i] = 2
			if i-2 >= 0 {
				dp[i] = dp[i-2] + 2
			}
		}

		// 2. s[i-1] 为 ')', 判断应该与s[i]配对的位置 s[i-dp[i-1]-1]是否可以成功配起来
		//    若可以成功配起来，则在原dp[i-1]的基础上，增加s[i]、s[i-dp[i-1]-1] 2个长度值以及 dp[i-dp[i-1]-2]对应的长度
		//    s[i]和s[i-dp[i-1]-1] 将dp[i-1] 和 dp[i-dp[i-1]-2] 连接起来
		//    如：")((...))"
		if s[i-1] == ')' && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
			dp[i] = dp[i-1] + 2
			if i-dp[i-1]-2 >= 0 {
				dp[i] = dp[i] + dp[i-dp[i-1]-2]
			}
		}

		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}
