package main

import (
	"strconv"
)

// 思路: 栈的应用

// 时间复杂度: O(n)    空间复杂度: O(n)

func evalRPN(tokens []string) int {
	n := len(tokens)

	stack := make([]string, 0)
	for i := 0; i < n; i++ {
		cur := tokens[i]
		if cur != "+" && cur != "-" && cur != "*" && cur != "/" {
			stack = append(stack, cur)
			continue
		}

		s1, s2 := stack[len(stack)-1], stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		num1, _ := strconv.Atoi(s1)
		num2, _ := strconv.Atoi(s2)

		val := 0
		if cur == "+" {
			val = num2 + num1
		}

		if cur == "-" {
			val = num2 - num1
		}

		if cur == "*" {
			val = num2 * num1
		}

		if cur == "/" {
			val = num2 / num1
		}

		sVal := strconv.Itoa(val)
		stack = append(stack, sVal)
	}

	res, _ := strconv.Atoi(stack[0])
	return res
}
