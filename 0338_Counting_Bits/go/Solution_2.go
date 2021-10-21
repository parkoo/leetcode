package main

// 位运算
// 时间复杂度：O(nlgn)  空间复杂度：O(1)

func countBits_1(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = getBits(i)
	}

	return res
}

// Brian Kernighan 算法原理：对于任意整数 x，令x=x&(x-1)，该运算将 x 的二进制表示的最后一个 1 变成 0.
// 因此，对 x 重复该操作，直到 x 变成 0，则操作次数即为 x 二进制表示中 1 的个数，时间复杂度为 lg(n).
func getBits(x int) int {
	cnt := 0
	for ; x != 0; x = x & (x - 1) {
		cnt++
	}

	return cnt
}
