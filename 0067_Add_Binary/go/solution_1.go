package main

import "strconv"

// 思路: 模拟法

// 时间复杂度: O(n)    空间复杂度: O(1)

func addBinary(a string, b string) string {
	res := ""

	carry := 0
	i, j := len(a)-1, len(b)-1
	for i >= 0 && j >= 0 {
		sum := int(a[i]-'0') + int(b[j]-'0') + carry
		cur := sum % 2
		carry = sum / 2
		res = strconv.Itoa(cur) + res
		i--
		j--
	}

	for i >= 0 {
		sum := int(a[i]-'0') + carry
		cur := sum % 2
		carry = sum / 2
		res = strconv.Itoa(cur) + res
		i--
	}

	for j >= 0 {
		sum := int(b[j]-'0') + carry
		cur := sum % 2
		carry = sum / 2
		res = strconv.Itoa(cur) + res
		j--
	}

	if carry != 0 {
		res = strconv.Itoa(carry) + res
	}

	return res
}
