package main

import (
	"strings"
)

// 思路: map的使用

// 时间复杂度：O(n), 空间复杂度：O(n)

func wordPattern_2(pattern string, s string) bool {
	ps := []rune(pattern)
	ss := strings.Split(s, " ")
	if len(ps) != len(ss) {
		return false
	}

	psMap := make(map[rune]string)
	sCache := make(map[string]struct{})

	for i, pVal := range ps {
		sVal := ss[i]
		if val, ok := psMap[pVal]; ok {
			if val != sVal {
				return false
			}

		} else {
			if _, ok := sCache[sVal]; ok {
				return false
			}

			psMap[pVal] = sVal
			sCache[sVal] = struct{}{}
		}
	}

	return true
}
