package main

// 当前数字大于右侧数字，加上当前数字；当前数字小于右侧数字，减去当前数字
// 时间复杂度：O(n)  空间复杂度：O(1)

func romanToInt_2(s string) int {
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
		if i < len(ss)-1 && romanMap[ss[i]] < romanMap[ss[i+1]] {
			res -= romanMap[ss[i]]
		} else {
			res += romanMap[ss[i]]
		}
	}

	return res
}
