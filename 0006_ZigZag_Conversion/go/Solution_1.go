package main

// 思路: 找规律 模拟法

// 时间复杂度: O(n)    空间复杂度: O(1)

func convert(s string, numRows int) string {
	res := make([]rune, 0)
	if numRows == 1 {
		return s
	}

	ss := []rune(s)

	// 一行一行处理
	for i := 0; i < numRows; i++ {
		downFlag := true

		if i >= len(ss) {
			break
		}

		cur := i
		for cur < len(ss) {
			res = append(res, ss[cur])

			// 计算steps (下一步的坐标)
			// numRows == 1 时， (numRows - (i+1)) * 2 ==  i * 2 == 0
			// 需要提前处理 numRows == 1 的情况

			// 第 0 行时：steps = (numRows - (i+1)) * 2
			// 第 numRows-1 行时：steps = i * 2
			// 其他行时：若向下走则 steps = (numRows - (i+1)) * 2，若向上走则 steps = i * 2, 同一行中，两者交替
			steps := (numRows - (i + 1)) * 2
			if !downFlag || i == numRows-1 {
				steps = i * 2
			}

			cur += steps

			// 非第0行和非第numRows-1行， 需要交替downFlag
			if i != 0 && i != numRows-1 {
				downFlag = !downFlag
			}
		}
	}

	return string(res)
}
