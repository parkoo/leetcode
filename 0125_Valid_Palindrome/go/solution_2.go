package main

import (
	"strings"
)

//思路： 双指针 对撞指针

// 时间复杂度：O(n)  空间复杂度：O(1) 原地处理

func isPalindrome_2(s string) bool {
	s = strings.ToLower(s)
	ss := []rune(s)

	isValidChar := func(a rune) bool {
		return (a >= '0' && a <= '9') || (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z')
	}

	// 双指针去除无效字符
	left := -1
	for right := 0; right < len(ss); right++ {
		if isValidChar(ss[right]) {
			left++
			ss[left], ss[right] = ss[right], ss[left]
		}
	}

	// 对撞指针做回文判断
	start, end := 0, left
	for start < end {
		if ss[start] != ss[end] {
			return false
		}
		start++
		end--
	}

	return true
}
