package main

// 思路: BFS

// 类似 lc0433

// 时间复杂度: O(mn)    空间复杂度: O(mn)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	q := make([]string, 0)
	q = append(q, beginWord)

	visited := make(map[string]bool)
	step := 0
	for len(q) > 0 {
		cnt := len(q)
		for i := 0; i < cnt; i++ {
			cur := q[0]
			q = q[1:]

			if cur == endWord {
				// step 是转换的次数，结果是需要包括起始单词的单词总长度，所以这里需要+1
				return step + 1
			}

			for _, next := range wordList {
				if canConvert(cur, next) && !visited[next] {
					q = append(q, next)
					visited[next] = true
				}
			}
		}
		step++
	}

	return 0
}

func canConvert(s, t string) bool {
	diffCnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			diffCnt++
		}
	}

	return diffCnt == 1
}
