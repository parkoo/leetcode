package main

import (
	"strings"
)

// map的使用
// 时间复杂度：O(n), 空间复杂度：O(n)

func wordPattern(pattern string, s string) bool {
	sArr := strings.Split(s, " ")
	if len(pattern) != len(sArr) {
		return false
	}

	m := make(map[string]string)
	for i := 0; i < len(sArr); i++ {
		v, ok := m[pattern[i:i+1]]
		if !ok {
			if inMapValue(m, sArr[i]) {
				return false
			}
			m[pattern[i:i+1]] = sArr[i]
		} else {
			if v != sArr[i] {
				return false
			}
		}
	}

	return true
}

func inMapValue(m map[string]string, s string) bool {
	for _, v := range m {
		if v == s {
			return true
		}
	}

	return false
}
