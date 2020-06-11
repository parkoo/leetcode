package main

// 常规暴力解法
// 时间复杂度：O(n^2)  空间复杂度：O(n)

func dailyTemperatures_1(T []int) []int {

	res := make([]int, len(T))

	for i:=0; i<len(T); i++ {
		for j:=i+1; j<len(T); j++ {
			if T[j]>T[i] {
				res[i] = j-i
				break
			}
		}
	}

	return res
}