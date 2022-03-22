package main

// 位运算  异或运算
// 异或运算性质： 1. X ^ 0 = X;   2. X ^ X = 0;   3. 满足交换律和结合律
// 时间复杂度：O(n)  空间复杂度；O(1)

func singleNumber_1(nums []int) int {
	temp := 0
	for _, num := range nums {
		temp ^= num
	}

	return temp
}
