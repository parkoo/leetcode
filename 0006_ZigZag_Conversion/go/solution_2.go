package main

// 思路: 找规律 模拟法 同solution_1

// 时间复杂度: O(n)    空间复杂度: O(1)

func convert_2(s string, numRows int) string {
	res := make([]rune, 0)
	if numRows == 1 {
		return s
	}

	ss := []rune(s)
	for i := 0; i < numRows; i++ {
		if i >= len(s) {
			break
		}

		downFlag := true
		cur := i
		for cur < len(s) {
			res = append(res, ss[cur])
			step := 0
			if i == 0 {
				step = (numRows - (i + 1)) * 2

			} else if i == numRows-1 {
				step = 2 * i

			} else {
				if downFlag {
					step = (numRows - (i + 1)) * 2
				} else {
					step = 2 * i
				}

				downFlag = !downFlag
			}

			cur += step
		}
	}

	return string(res)
}
