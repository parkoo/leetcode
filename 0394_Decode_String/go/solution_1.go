package main

// 思路: 通过辅助栈实现，注意复合的情况 3[a20[b]]

// 时间复杂度: O(s)    空间复杂度: O(s)  其中，s为解码后得出的字符串长度

func decodeString(s string) string {
	ss := []rune(s)

	stack := make([]rune, 0)
	for _, letter := range ss {
		// 对于非']'的字符，直接入栈
		if letter != ']' {
			stack = append(stack, letter)
			continue
		}

		// 遇到']'时，需要处理之前入栈的字符
		// 1，处理非数字字符，即字母或者'['
		subStr := make([]rune, 0)
		for len(stack) > 0 && (stack[len(stack)-1] < '0' || stack[len(stack)-1] > '9') {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != '[' {
				subStr = append(subStr, top)
			}
		}

		// 2，处理数字字符，计算出倍数
		times, factor := 0, 1
		for len(stack) > 0 && (stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9') {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			times = int(top-'0')*factor + times
			factor *= 10
		}

		// 3，将栈中取出的字符按倍数再放回栈中
		for i := 0; i < times; i++ {
			for j := len(subStr) - 1; j >= 0; j-- { // 这里保持之前入栈的顺序，需要倒着取值
				stack = append(stack, subStr[j])
			}
		}
	}

	// 将栈中的字符倒着取出即为结果
	cnt := len(stack)
	res := make([]rune, cnt)
	for i := cnt - 1; i >= 0; i-- {
		letter := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res[i] = letter
	}

	return string(res)
}
