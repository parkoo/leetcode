package main

// 时间复杂度：O(n)  空间复杂度：O(n)

func reverseWords(s string) string {
	// 去除头尾空格
	start, end := 0, len(s)-1
	for s[start] == ' ' {
		start++
	}
	for s[end] == ' ' {
		end--
	}

	// 空格去重 规范中间部分空格
	ss := make([]rune, 0)
	for i := start; i <= end; i++ {
		if s[i] != ' ' || (len(ss) > 0 && ss[len(ss)-1] != ' ') {
			ss = append(ss, rune(s[i]))
		}
	}

	i, j := 0, 0
	for ; j < len(ss); j++ {
		if ss[j] == ' ' || j == len(ss)-1 {
			k := j - 1
			if j == len(ss)-1 {
				k = j
			}
			reverse(ss, i, k)
			i = j + 1
		}
	}

	reverse(ss, 0, len(ss)-1)
	return string(ss)
}

func reverse(ss []rune, i, j int) {
	for i < j {
		ss[i], ss[j] = ss[j], ss[i]
		i++
		j--
	}
}
