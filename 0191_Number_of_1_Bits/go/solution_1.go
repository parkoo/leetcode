package main

// 思路: 汉明距离

// 时间复杂度: O(logn) 循环次数等于 n 的二进制位中 1 的个数，最坏情况下 n 的二进制位全部为 1, 需要循环 logn 次

// 空间复杂度: O(1)

func hammingWeight(n int) int {
	res := 0

	for n > 0 {
		n = n & (n - 1)
		res++
	}

	return res
}
