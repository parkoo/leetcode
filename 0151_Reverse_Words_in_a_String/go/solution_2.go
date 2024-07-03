package main

// 思路： 去除无效空格，整体翻转 再对每个单词翻转
// 使用双指针在原地去除无效空格 无额外空间

// 时间复杂度：O(n)  空间复杂度：O(1)

func reverseWords_2(s string) string {
	ss := []rune(s)

	// 双指针 移除无效空格
	left := -1
	spaceCnt := 0
	for right := 0; right < len(ss); right++ {
		cur := ss[right]
		if cur == ' ' {
			spaceCnt++
		} else {
			spaceCnt = 0
		}

		// ‘非空格’ || ‘非首字母且是字符后的第一个空格’ 为有效字符
		if (cur != ' ') || (cur == ' ' && spaceCnt == 1 && left != -1) {
			left++
			ss[left], ss[right] = ss[right], ss[left]
		}
	}
	if ss[left] == ' ' {
		left--
	}
	ss = ss[:left+1]

	// 整体翻转
	reverse(ss, 0, len(ss)-1)

	// 对每个单词翻转
	start := 0
	for end := 0; end <= len(ss)-1; end++ {
		if ss[end] == ' ' {
			reverse(ss, start, end-1)

			start = end + 1
			continue
		}

		if end == len(ss)-1 {
			reverse(ss, start, end)
		}
	}

	return string(ss)
}

func reverse(arr []rune, i, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
