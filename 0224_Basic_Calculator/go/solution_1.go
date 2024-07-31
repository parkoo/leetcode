package main

import (
	"strings"
)

// 思路:  通过两个栈来计算中缀表达式的值
// 比较 lc0150 同过一个栈来计算后缀表达式

// 时间复杂度: O(n)    空间复杂度: O(n)

func calculate(s string) int {
	s = prepare(s)
	ss := []rune(s)
	n := len(ss)

	ops, nums := make([]rune, 0), make([]int, 0)

	for i := 0; i < n; i++ {
		c := ss[i]

		// '(' 无脑入ops栈
		if c == '(' {
			ops = append(ops, c)

			// ')' 需要做计算,直到遇到与之匹配的'('
		} else if c == ')' {
			for len(ops) > 0 {
				curOp := ops[len(ops)-1]
				ops = ops[:len(ops)-1]

				// 遇到'('时， 出栈并结束循环
				if curOp == '(' {
					break
				}

				// 从nums中取值、计算、结果放回nums
				b := nums[len(nums)-1]
				nums = nums[:len(nums)-1]
				a := nums[len(nums)-1]
				nums = nums[:len(nums)-1]
				ans := calc(a, b, curOp)
				nums = append(nums, ans)
			}

		} else {
			// 数字, 需处理连续数字
			if isNum(c) {
				theNum := 0
				k := i
				for ; k < n && isNum(ss[k]); k++ {
					theNum = theNum*10 + int(ss[k]-'0')
				}
				nums = append(nums, theNum)
				i = k - 1

				// 各种操作符, 需要保证ops入栈时，栈顶的操作符优先级要小于当前待入栈的操作符优先级
				// 否则需要需要先将栈顶操作符出栈并做计算
			} else {
				for len(ops) > 0 && ops[len(ops)-1] != '(' &&
					priority(ops[len(ops)-1]) >= priority(c) {
					curOp := ops[len(ops)-1]
					ops = ops[:len(ops)-1]
					b := nums[len(nums)-1]
					nums = nums[:len(nums)-1]
					a := nums[len(nums)-1]
					nums = nums[:len(nums)-1]
					ans := calc(a, b, curOp)
					nums = append(nums, ans)
				}
				ops = append(ops, c)
			}

		}
	}

	// 处理ops栈中最后剩余的操作
	for len(ops) > 0 && ops[len(ops)-1] != '(' {
		curOp := ops[len(ops)-1]
		ops = ops[:len(ops)-1]
		b := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		a := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		ans := calc(a, b, curOp)
		nums = append(nums, ans)
	}

	return nums[0]
}

// 操作符的优先级
func priority(op rune) int {
	m := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	return m[op]
}

// 预处理
func prepare(s string) string {
	// 去除空格
	s = strings.ReplaceAll(s, " ", "")

	// (-1) -> (0-1)
	// (+1) -> (0+1)  本题不存在
	s = strings.ReplaceAll(s, "(-", "(0-")

	// 额外处理第一个非空字符为符号的情况  -23 -> 0-23
	if s[0] == '-' {
		s = "0" + s
	}

	return s
}

// 判断是否为数字
func isNum(c rune) bool {
	return int(c-'0') >= 0 && int(c-'0') <= 9
}

func calc(a, b int, op rune) int {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}
	return 0
}
