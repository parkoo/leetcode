package main

// 思路: 依次计算二进制的每一位数,
// 答案的第 i 个二进制位就是数组中所有元素的第 i 个二进制位之和除以 3 的余数

// 时间复杂度: O(nlogC)    空间复杂度: O(1)

func singleNumber_2(nums []int) int {
	res := int32(0) // 注意这里使用 int32

	for i := 0; i < 32; i++ {
		sum := int32(0) // 注意这里使用 int32
		for _, num := range nums {
			sum += int32(num) >> i & 1
		}

		if sum%3 == 1 {
			res = res | 1<<i
		}
	}

	return int(res)
}
