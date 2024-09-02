package main

// 思路: 计算从0位置开始的最长回文串，将剩余的部分翻转后拼接到头部
// 非最优解

// 时间复杂度: O(n^2)    空间复杂度: O(1)

func shortestPalindrome_1(s string) string {
	n := len(s)
	if n == 1 {
		return s
	}

	// 从后向前寻找以s[0]开头的最长回文串
	i := n - 1
	for i >= 0 {
		if isPalindrome(s, 0, i) {
			break
		}
		i--
	}

	// s 本身就是回文串
	if i == n-1 {
		return s
	}

	// 翻转剩余部分 拼接到头部
	sub := reverse(s, i+1, n-1)
	res := sub + s
	return res
}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func reverse(s string, l, r int) string {
	ss := []rune(s)
	i, j := l, r
	for i < j {
		ss[i], ss[j] = ss[j], ss[i]
		i++
		j--
	}

	return string(ss[l : r+1])
}
