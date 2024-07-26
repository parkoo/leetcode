package main

// 思路: 模拟加法

// 时间复杂度: O(n)    空间复杂度: O(1)

func plusOne(digits []int) []int {
	res := make([]int, 0)

	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		if i == len(digits)-1 {
			sum += 1
		}

		sub := sum % 10
		carry = sum / 10

		res = append(res, sub)
	}

	if carry != 0 {
		res = append(res, carry)
	}

	// 反转res
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}
