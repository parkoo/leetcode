package main

// 反转一半的数字
// 时间复杂度：O(log10(n))　　空间复杂度：O(1)

func isPalindrome_2(x int) bool {

	if x<0 || (x%10==0 && x!=0) {
		return false
	}
	
	y := 0
	for x>y {
		y = 10*y + x%10
		x /= 10
	}

	return x==y || x==y/10
}