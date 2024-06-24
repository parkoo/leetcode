package main

// 思路: 使用辅助栈

// 时间复杂度: O(n)    空间复杂度: O(n)

func longestValidParentheses(s string) int {
	res := 0

	stack := make([]int, 0)   // 栈底始终维护上一个不合法的‘)’的索引,
	stack = append(stack, -1) // 设置一个预处理值 -1, 必须是-1, 考虑"()"的情况

	for i := 0; i < len(s); i++ {
		// 遇到'(',直接将其索引入栈
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}

		// 遇到')', 先出栈，出栈后若栈为空，则说明该')'左边无相应的'('匹配，将该')'的索引入栈，即为当前最后一个无效的')'
		// 若出栈后栈不为空，则说明该')'左边有相应的'('匹配，此时以该')'为结尾的有效括号的长度为 i-stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(stack) == 0 {
			stack = append(stack, i)

		} else {
			// 当前元素位置与栈顶元素位置的长度,栈顶元素有可能是上一个无效的')'的位置
			// 可以考虑")()()"、")(())"两种情况
			curLen := i - stack[len(stack)-1]
			if curLen > res {
				res = curLen
			}
		}
	}

	return res
}
