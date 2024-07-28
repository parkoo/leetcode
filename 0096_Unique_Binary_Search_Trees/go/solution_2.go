package main

// 思路: 卡塔兰数公式
// C(0) = 1,  C(n+1) = C(n) * 2(2n+1) / (n+2)

// 时间复杂度：O(n^2)  空间复杂度：O(n)

func numTrees_2(n int) int {

	res := 1
	for i := 0; i < n; i++ {
		res = 2 * (2*i + 1) * res / (i + 2)
	}

	return res
}
