package main

// 思路：模拟法，罗马数字是由固定精度相加而得的，每次直接减去比当前值小的最大的精度即可

// 时间复杂度: O(1)    空间复杂度: O(1)

func intToRoman(num int) string {
	// 精度从大到小排列
	valueSymbols := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	res := ""
	for num > 0 {
		// 寻找比当前num小的最大的value
		value, symbol := 0, ""
		for _, valueSymbol := range valueSymbols {
			if num >= valueSymbol.value {
				value, symbol = valueSymbol.value, valueSymbol.symbol
				break
			}
		}

		num -= value
		res += symbol
	}

	return res
}
