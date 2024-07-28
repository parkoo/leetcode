package main

// 思路: 使用map作记忆化查找

// 时间复杂度: O(n)    空间复杂度: O(n)

func longestSubsequence_2(arr []int, difference int) int {
	res := 0

	// cache[k] = u, 表示以数字k结尾的公差为difference的等差子序列的最长长度为u
	cache := make(map[int]int)
	for _, v := range arr {
		cache[v] = cache[v-difference] + 1
		if cache[v] > res {
			res = cache[v]
		}
	}

	return res
}
