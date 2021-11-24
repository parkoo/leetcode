package main

// 动态规划
// 时间复杂度：O(n)  空间复杂度：O(1)

func fib_1(n int) int {
	if n < 2 {
		return n
	}

	res := 0
	pre1, pre2 := 0, 1
	for i := 2; i <= n; i++ {
		res = pre1 + pre2
		pre1 = pre2
		pre2 = res
	}
	return res
}
