package main

// 单调栈
// 时间复杂度: O(mn)  空间复杂度: O(mn)

func maximalRectangle(matrix [][]byte) int {
	res := 0
	heights := make([]int, len(matrix[0])+2)

	for _, row := range matrix {
		for i, cell := range row {
			if cell == '0' { // 下层cell为‘0’时，柱高置0
				heights[i+1] = 0
			} else {
				heights[i+1] += 1
			}
		}

		stack := make([]int, 0)
		for i, h := range heights {
			for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
				index := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				curArea := (i - stack[len(stack)-1] - 1) * heights[index]
				if curArea > res {
					res = curArea
				}
			}
			stack = append(stack, i)
		}
	}

	return res
}
