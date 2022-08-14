package main

// 模拟/动态规划
// 时间复杂度：O(n)  空间复杂度：O(1)

func maxScore_3(s string) int {
	res := 0
	if len(s) == 0 {
		return res
	}

	for i := 1; i < len(s); i++ {
		res += int(s[i] - '0')
	}
	if int(s[0]-'0') == 0 {
		res++
	}

	// 在当前状态下， 将第i个元素划分到左边的数组，若s[i]='0', 则左边的数组值加一，右边的数组值不变，整体值加一，反之，整体值减一
	pre := res
	for i := 1; i < len(s)-1; i++ {
		if int(s[i]-'0') == 0 {
			pre++
		} else {
			pre--
		}

		if pre > res {
			res = pre
		}
	}

	return res
}
