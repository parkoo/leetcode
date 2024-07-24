package main

// 思路:  二分法 快速幂

// 时间复杂度: O(logn)    空间复杂度: O(logn) 递归栈

func myPow_1(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / quickPow(x, -n)
	}

	return quickPow(x, n)
}

func quickPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	y := quickPow(x, n/2)

	curVal := y * y
	if n%2 == 1 {
		curVal *= x
	}

	return curVal
}
