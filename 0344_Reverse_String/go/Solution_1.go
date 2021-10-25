package main

// 双指针 对撞指针
// 时间复杂度：O(n)  空间复杂度：O(1)

func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
