package main

// 思路: 双指针

// 时间复杂度: O(m+n)    空间复杂度: O(1)

func isSubsequence_1(s string, t string) bool {
	if len(s) > len(t) || (len(s) == len(t) && s != t) {
		return false
	}

	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == len(s)
}
