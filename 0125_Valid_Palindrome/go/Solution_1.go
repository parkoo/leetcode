package main

import (
	"strings"
)

// 双指针 对撞指针
// 时间复杂度：O(n)  空间复杂度：O(1)

func isPalindrome(s string) bool {
	ss := strings.ToUpper(s)
	sb := strings.Builder{}
	for _, v := range []rune(ss) {
		if (v >= '0' && v <= '9') || (v >= 'A' && v <= 'Z') {
			sb.WriteRune(v)
		}
	}
	ss = sb.String()

	i, j := 0, len(ss)-1
	for i < j {
		if ss[i:i+1] != ss[j:j+1] {
			return false
		}
		i++
		j--
	}

	return true
}
