package main

// 思路: 频率统计 同lc0242

// 时间复杂度: O(m+n)    空间复杂度: O(n)

func canConstruct_1(ransomNote string, magazine string) bool {
	ss, tt := []rune(ransomNote), []rune(magazine)
	tFreq := make(map[rune]int)
	for i := 0; i < len(tt); i++ {
		tFreq[tt[i]]++
	}

	for _, c := range ss {
		if cnt, ok := tFreq[c]; !ok || cnt == 0 {
			return false
		}

		tFreq[c]--
	}

	return true
}
