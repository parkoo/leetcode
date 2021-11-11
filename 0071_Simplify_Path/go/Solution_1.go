package main

import (
	"strings"
)

// 栈的应用
// 时间复杂度：O(n)  空间复杂度：O(n)

func simplifyPath(path string) string {
	ss := strings.Split(path, "/")
	stack := make([]string, 0)
	for i := 0; i < len(ss); i++ {
		if ss[i] == "" || ss[i] == "." {
			continue
		}

		if ss[i] == ".." {
			// 出栈判断
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}

		stack = append(stack, ss[i])
	}

	if len(stack) == 0 {
		return "/"
	}

	sb := &strings.Builder{}
	for i := 0; i < len(stack); i++ {
		sb.WriteString("/")
		sb.WriteString(stack[i])
	}

	return sb.String()
}
