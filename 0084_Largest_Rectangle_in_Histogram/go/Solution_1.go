package main

// 思路: 单调栈　维护一个单调递增栈，栈内存储索引

// 时间复杂度：O(n)  空间复杂度：O(n)

func largestRectangleArea_1(heights []int) int {
	res := 0

	q := []int{} // 通过slice维护一个单调递增栈
	for i := 0; i < len(heights); i++ {
		for len(q) > 0 && heights[i] < heights[q[len(q)-1]] {
			h := heights[q[len(q)-1]] // 以出栈元素为高，计算最大矩形的面积
			q = q[:len(q)-1]

			var w int // 计算宽
			if len(q) == 0 {
				w = i
			} else {
				w = i - q[len(q)-1] - 1
			}

			s := h * w
			if s > res {
				res = s
			}
		}

		q = append(q, i)
	}

	// 清空栈内元素,确保以每个元素作为高，并计算其面积
	for len(q) > 0 {
		h := heights[q[len(q)-1]] // 以出栈元素为高，计算最大矩形的面积
		q = q[:len(q)-1]

		var w int // 计算宽
		if len(q) > 0 {
			w = len(heights) - q[len(q)-1] - 1
		} else {
			w = len(heights)
		}

		s := h * w
		if s > res {
			res = s
		}
	}

	return res
}
