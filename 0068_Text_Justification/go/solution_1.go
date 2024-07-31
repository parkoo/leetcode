package main

import "strings"

// 思路: 模拟法 双指针

// 根据题目中填充空格的细节，分以下三种情况讨论
// 1. 当前行是最后一行: 单词左对齐，且单词之间应只有一个空格，在行末填充剩余空格
// 2. 当前行不是最后一行，且只有一个单词:该单词左对齐，在行末填充空格
// 3. 当前行不是最后一行，且不只一个单词: 需要处理单词间的空格分配

// 时间复杂度: O(m)    空间复杂度: O(m)

func fullJustify(words []string, maxWidth int) []string {
	res := make([]string, 0)

	n := len(words)

	left, right := 0, 0
	curLen := len(words[0])
	for right < n {
		// 向后寻找可以放入一行内的单词
		for right+1 < n && curLen+len(words[right+1])+1 <= maxWidth {
			right++
			curLen = curLen + len(words[right]) + 1
		}

		// 1. right到达最后一个单词，说明是到了最后一行
		if right == n-1 {
			item := strings.Join(words[left:], " ")
			tail := blank(maxWidth - len(item))
			item += tail
			res = append(res, item)
			break // 退出循环
		}

		// 2. 非最后一行, 且此行只有一个单词
		if right == left {
			item := words[left] + blank(maxWidth-len(words[left]))
			res = append(res, item)

			// 3. 非最后一行, 且此行有大于一个单词
		} else {
			wordLen := 0 // 统计该行内非空格的长度
			for i := left; i <= right; i++ {
				wordLen += len(words[i])
			}
			spaceLen := maxWidth - wordLen // 该行内空格的长度

			intervalCnt := right - left
			avgBlank := spaceLen / (intervalCnt)   // 平均每个单词之间的空格数
			extraBlank := spaceLen % (intervalCnt) // 多出来的空格数，需再均匀分配给前边的间隔内
			item := ""
			for i := left; i <= right; i++ {
				item += words[i]
				if i != right && extraBlank > 0 {
					item += blank(avgBlank + 1)
					extraBlank--

				} else if i != right && extraBlank == 0 {
					item += blank(avgBlank)
				}
			}
			res = append(res, item)
		}

		// 向下迭代
		right++
		left = right
		curLen = len(words[left])
	}

	return res
}

// 返回由n个空格组成的字符串
func blank(n int) string {
	return strings.Repeat(" ", n)
}
