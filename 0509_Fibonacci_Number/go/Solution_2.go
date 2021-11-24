package main

// 递归法
// 时间复杂度：O(n^2)   空间复杂度：O(n^2) 函数栈带来的开销

func fib_2(n int) int {
	if n < 2 {
		return n
	}

	return fib_2(n-1) + fib_2(n-2)
}
