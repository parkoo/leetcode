package main

import (
	"fmt"
)

// 回溯法
// 左括号数小于n时可以添加左括号 右括号数小于左括号数时可以添加右括号
// 时间复杂度：O[4n/n^(1/2)]  空间复杂度：O(n)

func generateParenthesis_1(n int) []string {
	res := make([]string, 0)

	var backtrack func(open, close, max int, subStr string)
	backtrack = func(open, close, max int, subStr string) {
		if len(subStr) == 2*max {
			res = append(res, subStr)
			return
		}

		// 左括号数小于n
		if open < max {
			subStr = fmt.Sprintf("%s%s", subStr, "(")
			backtrack(open+1, close, max, subStr)
			subStr = subStr[:len(subStr)-1]
		}

		// 右括号小于左括号数
		if close < open {
			subStr = fmt.Sprintf("%s%s", subStr, ")")
			backtrack(open, close+1, max, subStr)
			subStr = subStr[:len(subStr)-1]
		}
	}

	backtrack(0, 0, n, "")
	return res
}
