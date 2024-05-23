package main

// 思路: 暴力解法（超时）

// 时间复杂度: O(?)    空间复杂度: O(?)

func groupAnagrams_1(strs []string) [][]string {
	res := make([][]string, 0)

	for _, str := range strs {
		flag := false
		for i, group := range res {
			if isAnagrams(group[0], str) {
				group = append(group, str)
				res[i] = group
				flag = true
				break
			}
		}
		if !flag {
			res = append(res, []string{str})
		}
	}

	return res
}

// 判断两个字符串是否是字母异位词
func isAnagrams(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aFrequent, bFrequent := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(a); i++ {
		aFrequent[a[i]]++
		bFrequent[b[i]]++
	}

	for letter, aCnt := range aFrequent {
		if bCnt, ok := bFrequent[letter]; !ok || bCnt != aCnt {
			return false
		}
	}

	return true
}
