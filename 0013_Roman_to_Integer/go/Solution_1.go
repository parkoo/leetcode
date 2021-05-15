package main

// 时间复杂度：O(n)  空间复杂度：O(1)

func romanToInt_1(s string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	ss := []rune(s)
	res := 0
	for i := 0; i < len(ss); i++ {
		if i+1 < len(ss) && (romanMap[ss[i+1]]-romanMap[ss[i]] == 4 || romanMap[ss[i+1]]-romanMap[ss[i]] == 9 ||
			romanMap[ss[i+1]]-romanMap[ss[i]] == 40 || romanMap[ss[i+1]]-romanMap[ss[i]] == 90 ||
			romanMap[ss[i+1]]-romanMap[ss[i]] == 400 || romanMap[ss[i+1]]-romanMap[ss[i]] == 900) {
			res += romanMap[ss[i+1]] - romanMap[ss[i]]
			i++
		} else {
			res += romanMap[ss[i]]
		}
	}

	return res
}
