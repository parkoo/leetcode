package main

// 思路:  频次统计 + 滑动窗口
// 类似lc0438， 本题是将字符的组合改成了字符串的组合
// 例外需要对s进行分组，分组方式有span种，要对每一种分组进行滑动窗口统计

// 时间复杂度: O(n*span)    空间复杂度: O(m*span)  n为s的长度，m为words的长度, span为words中每个单词的长度

func findSubstring(s string, words []string) []int {
	res := make([]int, 0)

	n := len(s)
	span := len(words[0])
	total := len(words) * span
	tFreq := make(map[string]int)
	for _, word := range words {
		tFreq[word]++
	}

	// 对s按照span的长度分组，共有span种分法，即起始位置分别从[0,span-1]开始
	for start := 0; start < span; start++ {

		// 对每一种分组进行滑动窗口处理
		sFreq := make(map[string]int)
		validLen := 0
		left := start
		for right := start; right <= n-span; right += span {
			cur := s[right : right+span]
			sFreq[cur]++
			if cnt, ok := tFreq[cur]; ok && sFreq[cur] <= cnt {
				validLen++
			}

			for validLen == len(words) {
				if right+span-left == total {
					res = append(res, left)
				}

				old := s[left : left+span]
				left = left + span
				sFreq[old]--
				if cnt, ok := tFreq[old]; ok && sFreq[old] < cnt {
					validLen--
				}
			}
		}
	}

	return res
}
