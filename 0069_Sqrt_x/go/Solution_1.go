package main

// 变种二分法
// 时间复杂度：O(lgn)  空间复杂度：O(1)

func mySqrt(x int) int {
	l, r := 0, x
	for l < r {
		mid := l + (r-l)/2
		if mid*mid >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l*l == x {
		return l
	}
	return l - 1
}
