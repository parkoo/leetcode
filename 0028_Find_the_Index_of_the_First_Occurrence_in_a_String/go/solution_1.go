package main

// 思路: 遍历

// 时间复杂度: O(mn)    空间复杂度: O(1)

func strStr_1(haystack string, needle string) int {
	s, t := []rune(haystack), []rune(needle)

	for i := 0; i < len(s); i++ {
		m, n := i, 0
		for ; n < len(t); n++ {
			if m >= len(s) || s[m] != t[n] {
				break
			}
			m++
		}
		if n == len(t) {
			return i
		}
	}

	return -1
}
