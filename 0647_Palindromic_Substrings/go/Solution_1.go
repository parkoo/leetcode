package main

// 中心扩展法
// 时间复杂度：O(n^2)  空间复杂度：O(1)

func countSubstrings(s string) int {
	res := 0

	var helper func(ss []rune, i, j int)
	helper = func(ss []rune, i, j int) {
		for i >= 0 && j < len(ss) && ss[i] == ss[j] {
			res++
			i--
			j++
		}
	}

	ss := []rune(s)
	for i := 0; i < len(ss); i++ {
		helper(ss, i, i)
		helper(ss, i, i+1)
	}

	return res
}
