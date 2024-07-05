package main

import (
	"strings"
)

// 思路: 将path以“/”进行切分，切分后的结果可能是"."、".."、""以及其他字符串
// 遍历切分后的数组，遇到其他字符串入栈，遇到"."、""则不需要处理，遇到".."需要尝试出栈以回到上层目录
// 最终将栈内从栈底到栈顶元素用"/"相拼接便是结果，注意栈栈为空的情况

// 时间复杂度：O(n)  空间复杂度：O(n)

func simplifyPath_2(path string) string {
	names := strings.Split(path, "/")

	stack := make([]string, 0)

	for _, name := range names {
		if name != "." && name != ".." && name != "" {
			stack = append(stack, name)
		}

		if name == ".." && len(stack) != 0 {
			stack = stack[:len(stack)-1]
		}
	}

	res := ""
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = "/" + cur + res
	}
	if res == "" {
		res = "/"
	}

	return res
}
