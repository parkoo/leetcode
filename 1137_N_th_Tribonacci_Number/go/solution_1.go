package main

// 思路: 动态规划

// 时间复杂度: O(n)    空间复杂度: O(1)

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n < 3 {
		return 1
	}

	pre1, pre2, pre3 := 0, 1, 1
	res := 0
	for i := 3; i <= n; i++ {
		res = pre1 + pre2 + pre3

		pre1 = pre2
		pre2 = pre3
		pre3 = res
	}

	return res
}
