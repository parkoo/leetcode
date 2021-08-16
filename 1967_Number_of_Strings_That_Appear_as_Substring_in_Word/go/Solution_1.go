package main

// 时间复杂度：O(n*sum(mi)), 其中 n=len(word), mi=len(pattern[i])
// 空间复杂度：O(1)

func numOfStrings_1(patterns []string, word string) int {
	res := 0
	for _, p := range patterns {
		if isSubstr(p, word) {
			res++
		}
	}

	return res
}

func isSubstr(pattern, word string) bool {
	wrs := []rune(word)
	prs := []rune(pattern)

	if len(prs) > len(wrs) {
		return false
	}
	for i := 0; i <= len(wrs)-len(prs); i++ {
		var flag = true
		for j := i; j < i+len(prs); j++ {
			if wrs[j] != prs[j-i] {
				flag = false
			}
		}
		if flag {
			return true
		}
	}

	return false
}
