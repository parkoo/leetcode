package main

// 思路: 左右括号计数, 空间复杂度可优化到常数级别

// 时间复杂度: O(n)    空间复杂度: O(1)

func longestValidParentheses(s string) int {
	res := 0

	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			curLen := 2 * left
			if curLen > res {
				res = curLen
			}
		}
		if right > left { // 遇到无效的')'
			left, right = 0, 0
		}
	}

	// 上述过程处理了left == right 和 right > left 的情况，对于 right > left 的情况会遗漏
	// 需要额外反向处理一次
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			right++
		} else {
			left++
		}

		if left == right {
			curLen := 2 * left
			if curLen > res {
				res = curLen
			}
		}
		if left > right { // 遇到无效的'('
			left, right = 0, 0
		}
	}

	return res
}
