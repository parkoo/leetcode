package main

import (
	"fmt"
)

// next permutatuon
// 时间复杂度：O(max(k, n))  空间复杂度：O(1)

func getPermutation(n int, k int) string {
	s := ""
	for i := 1; i <= n; i++ {
		s = fmt.Sprintf("%s%d", s, i)
	}
	for i := 0; i < k-1; i++ {
		s = nextPermutation(s)
	}

	return s
}

func nextPermutation(s string) string {
	ss := []rune(s)
	n := len(ss)

	for i := n - 1; i > 0; i-- {
		// 从后向前寻找顺序对
		if ss[i-1] < ss[i] {
			// 交换
			for j := n - 1; j >= i; j-- {
				if ss[j] > ss[i-1] {
					ss[j], ss[i-1] = ss[i-1], ss[j]
					break
				}
			}

			// 翻转
			reverse(ss, i, n-1)
			break
		}
	}

	return string(ss)
}

func reverse(ss []rune, i, j int) {
	for i < j {
		ss[i], ss[j] = ss[j], ss[i]
		i++
		j--
	}
}
