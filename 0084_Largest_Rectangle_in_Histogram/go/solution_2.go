package main

// 思路: 同solution_1, 单调栈　维护一个单调递增栈，栈内存储索引

// 时间复杂度：O(n)  空间复杂度：O(n)

func largestRectangleArea_2(heights []int) int {
	res := 0

	stack := make([]int, 0) // 单调递增栈
	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			curIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 确定左右边界, 即左边第一个比出栈元素小的元素的索引和右边边第一个比出栈元素小的元素的索引
			// 再这两个边界之间，出栈元素为最小值，可以以出栈元素为高计算矩形面积
			left, right := curIndex, i
			weight := i // 若栈内没有元素，说明该出栈元素左侧的值均大于该元素 right - 0 = i
			if len(stack) > 0 {
				left = stack[len(stack)-1]
				weight = right - left - 1
			}

			area := weight * heights[curIndex]
			if area > res {
				res = area
			}
		}

		stack = append(stack, i)
	}

	// 清空栈内剩余元素，确保以所有元素为高的面积都被计算过
	// 栈内剩余元素出栈时，该元素右侧所有的值均大于该元素，即该元素右侧没有比该元素更大的值, 所以矩形的右边界right可以取len(heights)
	// 若元素出栈后，栈内还有元素，则该元素左侧第一个小于该元素的值即为此时的栈顶元素，此时，矩形左边界left取栈顶值stack[len(stack)-1]
	// 若元素出栈后，栈内无元素，则说明该元素左边没有比该元素更小的值，矩形左边界left取0，且左边界包含在内
	for len(stack) > 0 {
		curIndex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		left, right := 0, len(heights)
		weight := len(heights) // right - 0 = en(heights)
		if len(stack) > 0 {
			left = stack[len(stack)-1]
			weight = right - left - 1
		}

		area := weight * heights[curIndex]
		if area > res {
			res = area
		}
	}

	return res
}
