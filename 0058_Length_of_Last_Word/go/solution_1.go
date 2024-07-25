package main

// 思路: 反向遍历

// 时间复杂度: O(n)    空间复杂度: O(1)

func lengthOfLastWord(s string) int {
	ss := []rune(s)

	end := len(ss) - 1
	for ss[end] == ' ' {
		end--
	}

	cnt := 0
	for start := end; start >= 0; start-- {
		if ss[start] == ' ' {
			break
		}
		cnt++
	}

	return cnt
}
