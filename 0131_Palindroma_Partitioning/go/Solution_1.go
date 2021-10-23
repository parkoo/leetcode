package main

// 回溯法

func partition(s string) [][]string {
	res := make([][]string, 0)

	// 判断回文串
	var isPalindrome func(s string) bool
	isPalindrome = func(s string) bool {
		for i, j := 0, len(s)-1; i < j; {
			if s[i:i+1] != s[j:j+1] {
				return false
			}
			i++
			j--
		}

		return true
	}

	var backtrack func(start int, s string, subStrArr []string)
	backtrack = func(start int, s string, subStrArr []string) {
		if start == len(s) {
			arr := make([]string, len(subStrArr))
			for i, v := range subStrArr {
				arr[i] = v
			}
			res = append(res, arr)
			return
		}

		for i := 1; i <= len(s); i++ {
			end := start + i
			if end > len(s) {
				break
			}

			str := s[start:end]
			if isPalindrome(str) {
				subStrArr = append(subStrArr, str)
				backtrack(end, s, subStrArr)
				subStrArr = subStrArr[:len(subStrArr)-1]
			}
		}
	}

	backtrack(0, s, []string{})
	return res
}
