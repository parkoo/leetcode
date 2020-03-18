package main

// 反转整个数字（反转过程中可能溢出）
// 时间复杂度：O(log10(n))　　空间复杂度：O(1)

func isPalindrome_1(x int) bool {

	if x<0 {
		return false
	}
	
	y := 0
	z := x
	for x>0 {
		y = 10*y + x%10
		x /= 10
	}

	return z==y
}